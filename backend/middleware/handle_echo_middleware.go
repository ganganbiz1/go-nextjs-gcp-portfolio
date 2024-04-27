package middleware

import (
	"github.com/labstack/echo/v4"
)

type HandleEchoMiddleware struct {
}

func NewHandleEchoMiddleware() *HandleEchoMiddleware {
	return &HandleEchoMiddleware{}
}

func (m *HandleEchoMiddleware) Handle() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err != nil {
				c.Error(err)
			}
			return nil
		}
	}
}
