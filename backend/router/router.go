package router

import (
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/handler"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/middleware"
	"github.com/labstack/echo/v4"
)

type Router struct {
	userAuthMiddleware *middleware.UserAuthMiddleware
	healthcheckHandler handler.IfHealthcheckHandler
	userHandler        handler.IfUserHandler
	articleHandler     handler.IfArticleHandler
}

func NewRouter(
	userAuthMiddleware *middleware.UserAuthMiddleware,
	healthcheckHandler handler.IfHealthcheckHandler,
	userHandler handler.IfUserHandler,
	articleHandler handler.IfArticleHandler,

) *Router {
	return &Router{
		userAuthMiddleware,
		healthcheckHandler,
		userHandler,
		articleHandler,
	}
}

func (r *Router) Apply(e *echo.Echo) {
	g := e.Group("")
	group(g, "", []echo.MiddlewareFunc{}, func(g *echo.Group) {
		g.GET("/healthcheck", r.healthcheckHandler.Healthcheck)
		group(g.Group("/users"), "", []echo.MiddlewareFunc{}, func(g *echo.Group) {
			g.POST("/public", r.userHandler.PublicSignup)
			g.POST("", r.userHandler.Signup)
			group(g.Group("", r.userAuthMiddleware.Auth), "", []echo.MiddlewareFunc{}, func(g *echo.Group) {
				// 認証が必要なもの
				g.GET("/:userId", r.userHandler.Search)
			})
		})
		group(g.Group("", r.userAuthMiddleware.Auth), "", []echo.MiddlewareFunc{}, func(g *echo.Group) {
			// 認証が必要なもの
			group(g.Group("/articles"), "", []echo.MiddlewareFunc{}, func(g *echo.Group) {
				g.POST("", r.articleHandler.Create)
				g.GET("", r.articleHandler.List)
				g.GET("/:articleId", r.articleHandler.Get)
				g.PUT("/:articleId", r.articleHandler.Update)
				g.DELETE("/:articleId", r.articleHandler.Delete)
			})
		})

	})
}

func group(g *echo.Group, path string, ms []echo.MiddlewareFunc, f func(*echo.Group)) {
	f(g.Group(path, ms...))
}
