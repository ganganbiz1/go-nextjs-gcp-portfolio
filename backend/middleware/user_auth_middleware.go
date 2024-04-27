package middleware

import (
	"github.com/labstack/echo/v4"
)

type UserAuthMiddleware struct {
}

func NewUserAuthMiddleware() *UserAuthMiddleware {
	return &UserAuthMiddleware{}
}

func (m *UserAuthMiddleware) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// NOTE: ポートフォリオ用に公開するので、常に同じユーザでログイン状態にする
		c.Set("userID", 1)
		return next(c)
	}
}
