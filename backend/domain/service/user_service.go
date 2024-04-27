package service

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/entity"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/externals/gcp"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/repository"
)

type IfUserService interface {
	Signup(ctx context.Context, e *entity.User, password string) error
	SocialSignup(ctx context.Context, e *entity.User, idToken string) error
	PublicSignup(ctx context.Context, e *entity.User) error
	Login(ctx context.Context, email, idToken string) (*entity.TokenInfo, error)
	Search(ctx context.Context, id int) (*entity.User, error)
	Copy(ctx context.Context, ids []int) error
}

type UserService struct {
	userRepo       repository.IfUserRepository
	firebaseClient gcp.IfFirebaseClient
}

func NewUserService(
	userRepo repository.IfUserRepository,
	firebaseClient gcp.IfFirebaseClient,
) IfUserService {
	return &UserService{
		userRepo:       userRepo,
		firebaseClient: firebaseClient,
	}
}

func (s *UserService) Signup(ctx context.Context, e *entity.User, password string) error {
	eCount, err := s.userRepo.CountByEmail(ctx, e.Email)
	if err != nil {
		return domain.HandleError(domain.ErrInternal, err)
	}

	if eCount != 0 {
		return domain.HandleError(domain.ErrAlreadyExist, errors.New("email is already exist"))
	}

	nCount, err := s.userRepo.CountByName(ctx, e.Name)
	if err != nil {
		return domain.HandleError(domain.ErrInternal, err)
	}

	if nCount != 0 {
		return domain.HandleError(domain.ErrAlreadyExist, errors.New("name is already exist"))
	}

	fu, err := s.firebaseClient.Signup(ctx, e.Email, password)
	if err != nil {
		return domain.HandleError(domain.ErrInternal, err)
	}

	if err = s.userRepo.Create(ctx,
		entity.NewUser(0,
			e.Email,
			e.Name,
			fu.FirebaseUID,
			fu.FirebaseProviderID,
		)); err != nil {
		return domain.HandleError(domain.ErrInternal, err)
	}

	return nil
}

func (s *UserService) SocialSignup(ctx context.Context, e *entity.User, idToken string) error {
	_, err := s.firebaseClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		return domain.HandleError(domain.ErrAuth, err)
	}

	eCount, err := s.userRepo.CountByEmail(ctx, e.Email)
	if err != nil {
		return domain.HandleError(domain.ErrInternal, err)
	}

	if eCount != 0 {
		return domain.HandleError(domain.ErrAlreadyExist, errors.New("email is already exist"))
	}

	nCount, err := s.userRepo.CountByName(ctx, e.Name)
	if err != nil {
		return domain.HandleError(domain.ErrInternal, err)
	}

	if nCount != 0 {
		return domain.HandleError(domain.ErrAlreadyExist, errors.New("name is already exist"))
	}

	if err = s.userRepo.Create(ctx, e); err != nil {
		return domain.HandleError(domain.ErrInternal, err)
	}

	return nil
}

func (s *UserService) PublicSignup(ctx context.Context, e *entity.User) error {
	nCount, err := s.userRepo.CountByName(ctx, e.Name)
	if err != nil {
		return domain.HandleError(domain.ErrInternal, err)
	}

	if nCount != 0 {
		return domain.HandleError(domain.ErrAlreadyExist, errors.New("name is already exist"))
	}

	// NOTE: ポートフォリオ用として公開するので、ユーザ名以外はダミー値
	if err = s.userRepo.Create(ctx,
		entity.NewUser(0,
			fmt.Sprintf("%s@test.com", e.Name),
			e.Name,
			fmt.Sprintf("%s_uid", e.Name),
			fmt.Sprintf("%s_provider", e.Name),
		)); err != nil {
		return domain.HandleError(domain.ErrInternal, err)
	}

	return nil
}

func (s *UserService) Login(ctx context.Context, email, idToken string) (*entity.TokenInfo, error) {
	c, err := s.userRepo.CountByEmail(ctx, email)
	if err != nil {
		return nil, domain.HandleError(domain.ErrInternal, err)
	}

	if c == 0 {
		return nil, domain.HandleError(domain.ErrNotFound, errors.New("user NotFound"))
	}

	return s.firebaseClient.VerifyIDToken(ctx, idToken)
}

func (s *UserService) Search(ctx context.Context, id int) (*entity.User, error) {
	e, err := s.userRepo.Get(ctx, id)
	if err != nil && err != domain.ErrNotFound {
		return nil, domain.HandleError(domain.ErrInternal, err)
	}
	if e == nil || err == domain.ErrNotFound {
		return nil, domain.HandleError(domain.ErrNotFound, domain.ErrNotFound)
	}

	return e, nil
}

// CopyメソッドはGoルーチンのサンプル用。パブリックリポジトリに移行するときは移植しない
func (s *UserService) Copy(ctx context.Context, ids []int) error {

	errQnum := 1

	var err error
	errChan := make(chan error, errQnum)
	userChan := make(chan *entity.User)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		select {
		case <-ctx.Done():
			err = context.Canceled
			cancel()
			return
		case err, ok := <-errChan:
			if ok {
				_ = domain.HandleError(domain.ErrInternal, err)
				cancel() // Cancel the context to stop all operations
				return
			}
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(userChan)
		s.userRepo.ListWithChan(ctx, ids, userChan, errChan)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.userRepo.CreateWithChan(ctx, userChan, errChan)
	}()

	wg.Wait()

	close(errChan)

	if err != nil {
		return domain.HandleError(domain.ErrInternal, err)
	}

	return nil
}
