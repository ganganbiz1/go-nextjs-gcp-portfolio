package repository

import (
	"context"
	"time"

	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/entity"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/repository"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/infra"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/infra/repository/model"
)

type UserRepository struct {
	baseRepo  BaseRepository
	replicaDB ReplicaDB
}

func NewUserRepository(
	baseRepo BaseRepository,
	replicaDB ReplicaDB,
) repository.IfUserRepository {
	return &UserRepository{
		baseRepo,
		replicaDB,
	}
}

func (r *UserRepository) Create(ctx context.Context, e *entity.User) error {
	t := time.Now()
	exec := r.baseRepo.getExec(ctx)
	if err := exec.Model(&model.User{}).
		Create(map[string]interface{}{
			"email":                e.Email,
			"name":                 e.Name,
			"firebase_uid":         e.FirebaseUID,
			"firebase_provider_id": e.FirebaseProviderID,
			"created_at":           t,
			"updated_at":           t,
		}).Error; err != nil {
		return infra.HandleError(err)
	}
	return nil
}

func (r *UserRepository) Get(ctx context.Context, id int) (*entity.User, error) {
	var m model.User
	err := r.replicaDB.WithContext(ctx).
		Where("id = ?", id).
		First(&m).
		Error

	if err != nil {
		return nil, infra.HandleError(err)
	}

	return m.ToEntity(), nil
}

func (r *UserRepository) CountByEmail(ctx context.Context, email string) (int, error) {
	exec := r.baseRepo.getExec(ctx)
	var count int64
	if err := exec.
		Model(&model.User{}).
		Where("email = ?", email).
		Count(&count).
		Error; err != nil {
		return -1, infra.HandleError(err)
	}

	return int(count), nil
}

func (r *UserRepository) CountByName(ctx context.Context, name string) (int, error) {
	exec := r.baseRepo.getExec(ctx)
	var count int64
	if err := exec.Model(&model.User{}).
		Where("name = ?", name).
		Count(&count).
		Error; err != nil {
		return -1, infra.HandleError(err)
	}

	return int(count), nil
}

func (r *UserRepository) ListWithChan(ctx context.Context,
	ids []int,
	userChan chan<- *entity.User,
	errChan chan<- error,
) {
	var ms []*model.User
	err := r.replicaDB.WithContext(ctx).
		Where("id IN ?", ids).
		Find(&ms).
		Error

	if err != nil {
		errChan <- infra.HandleError(err)
		return
	}

	for _, m := range ms {
		userChan <- m.ToEntity()
	}
}

func (r *UserRepository) CreateWithChan(
	ctx context.Context,
	userChan <-chan *entity.User,
	errChan chan<- error,
) {
	for u := range userChan { // チャンネルが閉じられるまでループ
		t := time.Now()
		exec := r.baseRepo.getExec(ctx)
		if err := exec.Model(&model.User{}).
			Create(map[string]interface{}{
				"email":                u.Email,
				"name":                 u.Name,
				"firebase_uid":         u.FirebaseUID,
				"firebase_provider_id": u.FirebaseProviderID,
				"created_at":           t,
				"updated_at":           t,
			}).Error; err != nil {
			errChan <- infra.HandleError(err)
			return
		}
	}
}
