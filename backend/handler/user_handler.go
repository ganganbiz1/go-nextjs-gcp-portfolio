package handler

import (
	"net/http"
	"strconv"

	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/handler/presenter"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/usecase"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/usecase/dto/input"
	"github.com/labstack/echo/v4"
)

type IfUserHandler interface {
	Signup(c echo.Context) error
	PublicSignup(c echo.Context) error
	Search(c echo.Context) error
}

type UserHandler struct {
	userUsecase usecase.IfUserUsecase
}

type SignupRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required,password"`
}

type SocialSignupRequest struct {
	Email              string `json:"email" validate:"required,email"`
	Name               string `json:"name" validate:"required"`
	IDToken            string `json:"idToken" validate:"required"`
	FirebaseUID        string `json:"firebaseUid" validate:"required"`
	FirebaseProviderID string `json:"firebaseProviderId" validate:"required"`
}

type PublicSignupRequest struct {
	Name string `json:"name" validate:"required"`
}

type LoginRequest struct {
	Email   string `json:"email" validate:"required,email"`
	IDToken string `json:"idToken" validate:"required"`
}

func NewUserHandler(userUsecase usecase.IfUserUsecase) IfUserHandler {
	return &UserHandler{
		userUsecase,
	}
}

func (h *UserHandler) Signup(c echo.Context) error {
	var req SignupRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := c.Validate(&req); err != nil {
		return handleError(domain.ErrBadRequest, err)
	}

	if err := h.userUsecase.Signup(
		c.Request().Context(),
		input.NewUser(
			0,
			req.Email,
			req.Name,
			req.Password,
			"",
			"",
			"",
		),
	); err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}

func (h *UserHandler) SocialSignup(c echo.Context) error {
	var req SocialSignupRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := c.Validate(&req); err != nil {
		return handleError(domain.ErrBadRequest, err)
	}

	if err := h.userUsecase.Signup(
		c.Request().Context(),
		input.NewUser(
			0,
			req.Email,
			req.Name,
			"",
			req.IDToken,
			req.FirebaseUID,
			req.FirebaseProviderID,
		),
	); err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}

func (h *UserHandler) PublicSignup(c echo.Context) error {
	var req PublicSignupRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := c.Validate(&req); err != nil {
		return handleError(domain.ErrBadRequest, err)
	}

	if err := h.userUsecase.PublicSignup(
		c.Request().Context(),
		input.NewUser(
			0,
			"",
			req.Name,
			"",
			"",
			"",
			"",
		),
	); err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}

func (h *UserHandler) Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := c.Validate(&req); err != nil {
		return handleError(domain.ErrBadRequest, err)
	}

	e, err := h.userUsecase.Login(c.Request().Context(), req.Email, req.IDToken)
	if err != nil {
		return err
	}
	return presenter.BuildLoginResponse(c, e)
}

func (h *UserHandler) Search(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return err
	}
	e, err := h.userUsecase.Search(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return presenter.BuildUserSearchResponse(c, e)
}
