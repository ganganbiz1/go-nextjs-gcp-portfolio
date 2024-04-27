package server

import (
	"fmt"

	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/config"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/externals/datadog"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/handler"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/handler/validate"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	ddEcho "gopkg.in/DataDog/dd-trace-go.v1/contrib/labstack/echo.v4"
)

type Server struct {
	Echo         *echo.Echo
	serverConfig *config.ServerConfig
}

func NewServer(
	serverConfig *config.ServerConfig,
	newrelicMiddleware *middleware.NewRelicMiddleware,
	corsMiddleware *middleware.CorsMiddleware,
	handleEchoMiddleware *middleware.HandleEchoMiddleware,
	errHandler *handler.ErrorHandler,
	validator *validate.CustomValidator,
	dd datadog.IfDatadogClient,
) *Server {
	e := echo.New()

	e.HTTPErrorHandler = errHandler.Handler
	e.Validator = validator

	e.Use(ddEcho.Middleware(ddEcho.WithServiceName(dd.ServiceName("echo"))))
	e.Use(ddEcho.Middleware())

	e.Use(newrelicMiddleware.Setting())
	e.Use(corsMiddleware.Setting())
	e.Use(handleEchoMiddleware.Handle())
	e.Use(echoMiddleware.Recover())
	return &Server{
		Echo:         e,
		serverConfig: serverConfig,
	}
}

func (s *Server) Start() error {
	return s.Echo.Start(fmt.Sprintf(":%d", s.serverConfig.Port))
}
