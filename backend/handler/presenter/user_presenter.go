package presenter

import (
	"net/http"

	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/entity"
	"github.com/labstack/echo/v4"
)

type User struct {
	ID                 int    `json:"id"`
	Email              string `json:"email"`
	Name               string `json:"name"`
	FirebaseUID        string `json:"firebaseUId"`
	FirebaseProviderID string `json:"firebaseProviderId"`
}

type TokenInfo struct {
	AuthTime int    `json:"authTime"`
	Issuer   string `json:"issuer"`
	Audience string `json:"audience"`
	Expires  int    `json:"expires"`
	IssuedAt int    `json:"issuedAt"`
	Subject  string `json:"subject"`
	UID      string `json:"uid"`
}

func BuildLoginResponse(c echo.Context, e *entity.TokenInfo) error {
	return BuildSuccessResponse(c, http.StatusOK, &TokenInfo{
		AuthTime: e.AuthTime,
		Issuer:   e.Issuer,
		Audience: e.Audience,
		Expires:  e.Expires,
		IssuedAt: e.IssuedAt,
		Subject:  e.Subject,
		UID:      e.UID,
	})
}

func BuildUserSearchResponse(c echo.Context, e *entity.User) error {
	return BuildSuccessResponse(c, http.StatusOK, &User{
		ID:                 e.ID,
		Email:              e.Email,
		Name:               e.Name,
		FirebaseUID:        e.FirebaseUID,
		FirebaseProviderID: e.FirebaseProviderID,
	})
}
