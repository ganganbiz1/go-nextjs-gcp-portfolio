//go:build wireinject

//go:generate wire gen $GOFILE

package wire

import (
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/config"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/service"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/handler"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/handler/validate"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/infra/externals/datadog"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/infra/externals/gcp"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/infra/externals/newrelic"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/infra/repository"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/middleware"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/router"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/server"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/usecase"

	"github.com/google/wire"
)

type DIManager struct {
	Server *server.Server
	Router *router.Router
}

func DI() (*DIManager, func(), error) {
	wire.Build(
		wire.Struct(new(DIManager), "*"),

		server.NewServer,

		router.NewRouter,

		config.NewServerConfig,
		config.NewPrimaryDBConfig,
		config.NewReplicaDBConfig,
		config.NewFirebaseConfig,
		config.NewNewRelicConfig,
		config.NewDatadogConfig,

		middleware.NewCorsMiddleware,
		middleware.NewUserAuthMiddleware,
		middleware.NewNewRelicMiddleware,
		middleware.NewHandleEchoMiddleware,

		handler.NewErrorHandler,
		handler.NewHealthcheckHandler,
		handler.NewUserHandler,
		handler.NewArticleHandler,

		validate.NewCustomValidator,

		usecase.NewUserUsecase,
		usecase.NewArticleUsecase,

		service.NewUserService,
		service.NewArticleService,

		repository.NewBaseRepository,
		repository.NewPrimaryDBConnect,
		repository.NewReplicaDBConnect,
		repository.NewUserRepository,
		repository.NewArticleRepository,

		gcp.NewFirebaseClient,

		newrelic.NewClient,

		datadog.NewClient,
	)
	return &DIManager{}, nil, nil
}
