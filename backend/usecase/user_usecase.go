package usecase

import (
	"context"

	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/entity"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/service"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/usecase/dto/input"
)

type IfUserUsecase interface {
	Signup(ctx context.Context, dto *input.User) error
	SocialSignup(ctx context.Context, dto *input.User) error
	PublicSignup(ctx context.Context, dto *input.User) error
	Login(ctx context.Context, email, idToken string) (*entity.TokenInfo, error)
	Search(ctx context.Context, id int) (*entity.User, error)
}

type UserUsecase struct {
	userService service.IfUserService
}

func NewUserUsecase(
	userService service.IfUserService,
) IfUserUsecase {
	return &UserUsecase{
		userService: userService,
	}
}

func (u *UserUsecase) Signup(ctx context.Context, dto *input.User) error {
	if err := u.userService.Signup(ctx, dto.ToEntity(), dto.Password); err != nil {
		return handleError(domain.ErrInternal, err)
	}
	return nil
}

func (u *UserUsecase) SocialSignup(ctx context.Context, dto *input.User) error {
	if err := u.userService.SocialSignup(ctx, dto.ToEntity(), dto.IDToken); err != nil {
		return handleError(domain.ErrInternal, err)
	}
	return nil
}

func (u *UserUsecase) PublicSignup(ctx context.Context, dto *input.User) error {
	if err := u.userService.PublicSignup(ctx, dto.ToEntity()); err != nil {
		return handleError(domain.ErrInternal, err)
	}
	return nil
}

func (u *UserUsecase) Login(ctx context.Context, email, idToken string) (*entity.TokenInfo, error) {
	return u.userService.Login(ctx, email, idToken)
}

func (u *UserUsecase) Search(ctx context.Context, id int) (*entity.User, error) {
	e, err := u.userService.Search(ctx, id)
	if err != nil {
		return nil, handleError(domain.ErrInternal, err)
	}
	return e, nil
}
