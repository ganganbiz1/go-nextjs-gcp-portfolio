package service

import (
	"context"
	"errors"

	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/entity"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/repository"
)

type IfArticleService interface {
	Create(ctx context.Context, e *entity.Article) error
	List(ctx context.Context, userID int) ([]*entity.Article, error)
	Get(ctx context.Context, id, userID int) (*entity.Article, error)
	Update(ctx context.Context, e *entity.Article) error
	Delete(ctx context.Context, id, userID int) error
}

type ArticleService struct {
	userRepo    repository.IfUserRepository
	articleRepo repository.IfArticleRepository
}

func NewArticleService(
	userRepo repository.IfUserRepository,
	articleRepo repository.IfArticleRepository,
) IfArticleService {
	return &ArticleService{
		userRepo:    userRepo,
		articleRepo: articleRepo,
	}
}

func (s *ArticleService) Create(ctx context.Context, e *entity.Article) error {
	u, err := s.userRepo.Get(ctx, e.UserID)
	if err != nil {
		return domain.HandleError(domain.ErrInternal, err)
	}
	if u == nil {
		return domain.HandleError(domain.ErrConflict, errors.New("user does not exist"))
	}
	if err = s.articleRepo.Create(ctx, e); err != nil {
		return domain.HandleError(domain.ErrInternal, err)
	}
	return nil
}

func (s *ArticleService) List(ctx context.Context, userID int) ([]*entity.Article, error) {
	return s.articleRepo.List(ctx, userID)
}

func (s *ArticleService) Get(ctx context.Context, id, userID int) (*entity.Article, error) {
	return s.articleRepo.Get(ctx, id, userID)
}

func (s *ArticleService) Update(ctx context.Context, e *entity.Article) error {
	u, err := s.userRepo.Get(ctx, e.UserID)
	if err != nil {
		return domain.HandleError(domain.ErrInternal, err)
	}
	if u == nil {
		return domain.HandleError(domain.ErrConflict, errors.New("user does not exist"))
	}

	a, err := s.articleRepo.Get(ctx, e.ID, e.UserID)
	if err != nil {
		return domain.HandleError(domain.ErrInternal, err)
	}
	if a == nil {
		return domain.HandleError(domain.ErrNotFound, errors.New("article does not exist"))
	}

	if err = s.articleRepo.Update(ctx, e); err != nil {
		return domain.HandleError(domain.ErrInternal, err)
	}
	return nil
}
func (s *ArticleService) Delete(ctx context.Context, id, userID int) error {
	return s.articleRepo.Delete(ctx, id, userID)
}
