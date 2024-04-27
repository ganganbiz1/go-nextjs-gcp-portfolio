package gcp

import (
	"context"

	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/entity"
)

type IfFirebaseClient interface {
	Signup(ctx context.Context, email, password string) (*entity.User, error)
	VerifyIDToken(ctx context.Context, idToken string) (*entity.TokenInfo, error)
	// ResetPassword(ctx context.Context, email string) error
	// DeleteUser(ctx context.Context, email string) error
}
