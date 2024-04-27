package repository

import (
	"context"
	"time"

	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/entity"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/repository"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/infra"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/infra/repository/model"
)

type ArticleRepository struct {
	baseRepo  BaseRepository
	replicaDB ReplicaDB
}

func NewArticleRepository(
	baseRepo BaseRepository,
	replicaDB ReplicaDB,
) repository.IfArticleRepository {
	return &ArticleRepository{
		baseRepo,
		replicaDB,
	}
}

func (r *ArticleRepository) Create(ctx context.Context, e *entity.Article) error {
	t := time.Now()
	exec := r.baseRepo.getExec(ctx)
	if err := exec.Model(&model.Article{}).
		Create(map[string]interface{}{
			"title":      e.Title,
			"content":    e.Content,
			"image_id":   e.ImageID,
			"user_id":    e.UserID,
			"created_at": t,
			"updated_at": t,
		}).Error; err != nil {
		return infra.HandleError(err)
	}
	return nil
}

func (r *ArticleRepository) List(ctx context.Context, userID int) ([]*entity.Article, error) {
	var ms []*model.Article
	err := r.replicaDB.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("id asc").
		Find(&ms).
		Error

	if err != nil {
		return nil, infra.HandleError(err)
	}

	return model.Articles(ms).ToEntity(), nil
}

func (r *ArticleRepository) Get(ctx context.Context, id, userID int) (*entity.Article, error) {
	var m model.Article
	err := r.replicaDB.WithContext(ctx).
		Where("id = ?", id).
		Where("user_id = ?", userID).
		First(&m).
		Error

	if err != nil {
		return nil, infra.HandleError(err)
	}

	return m.ToEntity(), nil
}

func (r *ArticleRepository) Update(ctx context.Context, e *entity.Article) error {
	t := time.Now()
	exec := r.baseRepo.getExec(ctx)
	if err := exec.
		Model(&model.Article{}).
		Where("id = ?", e.ID).
		Where("user_id = ?", e.UserID).
		Updates(map[string]interface{}{
			"title":      e.Title,
			"content":    e.Content,
			"image_id":   e.ImageID,
			"user_id":    e.UserID,
			"updated_at": t,
		}).Error; err != nil {
		return infra.HandleError(err)
	}
	return nil
}

func (r *ArticleRepository) Delete(ctx context.Context, id, userID int) error {
	exec := r.baseRepo.getExec(ctx)
	if err := exec.
		Model(&model.Article{}).
		Where("id = ?", id).
		Where("user_id = ?", userID).
		Delete(&model.Article{}).
		Error; err != nil {
		return infra.HandleError(err)
	}
	return nil
}
