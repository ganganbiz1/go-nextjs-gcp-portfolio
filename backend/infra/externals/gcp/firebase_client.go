package gcp

import (
	"context"
	"errors"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/config"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/entity"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/externals/datadog"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/externals/gcp"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/infra"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/infra/externals/model"
	"google.golang.org/api/option"
)

type FirebaseClient struct {
	authClient *auth.Client
	dd         datadog.IfDatadogClient
}

func NewFirebaseClient(
	c *config.FirebaseConfig,
	dd datadog.IfDatadogClient,
) (gcp.IfFirebaseClient, error) {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsJSON([]byte(c.CredentialsJson)))
	if err != nil {
		return nil, err
	}
	authClient, err := app.Auth(ctx)
	if err != nil {
		return nil, infra.HandleError(fmt.Errorf("error getting Auth client: %v", err))

	}

	return &FirebaseClient{
		authClient: authClient,
		dd:         dd,
	}, nil
}

func (c *FirebaseClient) Signup(ctx context.Context, email, password string) (*entity.User, error) {
	params := (&auth.UserToCreate{}).
		Email(email).
		EmailVerified(false). // falseがデフォルト値
		Password(password).
		Disabled(false) // falseがデフォルト値

	// Datadog
	span, ctx := c.dd.StartSpan(ctx, "fibase-auth")
	defer span.Finish()

	u, err := c.authClient.CreateUser(ctx, params)
	if err != nil {
		return nil, infra.HandleError(errors.New("create user failed"))
	}

	return model.NewFirebaseUserInfo(email, u.UID, u.ProviderID).ToUserEntity(), nil
}

func (c *FirebaseClient) VerifyIDToken(ctx context.Context, idToken string) (*entity.TokenInfo, error) {
	// Datadog
	span, ctx := c.dd.StartSpan(ctx, "fibase-auth")
	defer span.Finish()

	t, err := c.authClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, infra.HandleError(errors.New("verify idToken failed"))
	}

	return entity.NewTokenInfo(
		int(t.AuthTime),
		t.Issuer,
		t.Audience,
		int(t.Expires),
		int(t.IssuedAt),
		t.Subject,
		t.UID,
	), nil
}
