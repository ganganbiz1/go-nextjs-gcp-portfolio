package middleware

import (
	myNewRelic "github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/infra/externals/newrelic"
	"github.com/labstack/echo/v4"
)

type NewRelicMiddleware struct {
	newRelic *myNewRelic.Client
}

func NewNewRelicMiddleware(
	newRelic *myNewRelic.Client,
) *NewRelicMiddleware {
	return &NewRelicMiddleware{
		newRelic,
	}
}

func (m *NewRelicMiddleware) Setting() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if m.newRelic.App != nil {
				txn := m.newRelic.App.StartTransaction(c.Request().URL.Path)
				defer txn.End()
			}
			return next(c)
		}
	}
}
