package usecase

import (
	"context"

	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/entity"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/service"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/usecase/dto/input"
)

type IfArticleUsecase interface {
	Create(ctx context.Context, dto *input.Article) error
	List(ctx context.Context, userID int) ([]*entity.Article, error)
	Get(ctx context.Context, id, userID int) (*entity.Article, error)
	Update(ctx context.Context, dto *input.Article) error
	Delete(ctx context.Context, id, userID int) error
}

type ArticleUsecase struct {
	articleService service.IfArticleService
}

func NewArticleUsecase(
	articleService service.IfArticleService,
) IfArticleUsecase {
	return &ArticleUsecase{
		articleService: articleService,
	}
}

func (u *ArticleUsecase) Create(ctx context.Context, dto *input.Article) error {
	return u.articleService.Create(ctx, dto.ToEntity())
}

func (u *ArticleUsecase) List(ctx context.Context, userID int) ([]*entity.Article, error) {
	return u.articleService.List(ctx, userID)
}

func (u *ArticleUsecase) Get(ctx context.Context, id, userID int) (*entity.Article, error) {
	return u.articleService.Get(ctx, id, userID)
}

func (u *ArticleUsecase) Update(ctx context.Context, dto *input.Article) error {
	return u.articleService.Update(ctx, dto.ToEntity())
}

func (u *ArticleUsecase) Delete(ctx context.Context, id, userID int) error {
	return u.articleService.Delete(ctx, id, userID)
}
