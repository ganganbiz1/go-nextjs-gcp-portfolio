package repository

import (
	"context"

	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/entity"
)

type IfUserRepository interface {
	Create(ctx context.Context, e *entity.User) error
	Get(ctx context.Context, id int) (*entity.User, error)
	CountByEmail(ctx context.Context, email string) (int, error)
	CountByName(ctx context.Context, name string) (int, error)
	ListWithChan(ctx context.Context, ids []int, userChan chan<- *entity.User, errChan chan<- error)
	CreateWithChan(ctx context.Context, userChan <-chan *entity.User, errChan chan<- error)
}
