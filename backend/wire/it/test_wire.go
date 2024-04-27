//go:build wireinject

//go:generate wire gen -output_file_prefix test_ .

package it

import (
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/handler"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/handler/validate"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/infra/repository"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/wire/it/di"
	"github.com/google/wire"
)

type DIManager struct {
	DIC *di.TestContainer
}

func DI() (*DIManager, func(), error) {
	wire.Build(
		wire.Struct(new(DIManager), "*"),

		di.NewTestContainer,

		di.NewTestServer,

		handler.NewErrorHandler,

		validate.NewCustomValidator,

		repository.NewBaseRepository,
		di.NewTestPrimaryDBConnect,
		di.NewTestReplicaDBConnect,
		repository.NewUserRepository,
	)
	return &DIManager{}, nil, nil
}
