package repository

import (
	"context"

	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/entity"
)

type IfArticleRepository interface {
	Create(ctx context.Context, e *entity.Article) error
	List(ctx context.Context, userID int) ([]*entity.Article, error)
	Get(ctx context.Context, id, userID int) (*entity.Article, error)
	Update(ctx context.Context, e *entity.Article) error
	Delete(ctx context.Context, id, userID int) error
}
