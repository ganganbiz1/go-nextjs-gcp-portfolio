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

type IfArticleHandler interface {
	Create(c echo.Context) error
	List(c echo.Context) error
	Get(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type ArticleHandler struct {
	articleUsecase usecase.IfArticleUsecase
}

type ArticleCreateRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
	ImageID int    `json:"imageId" validate:"required"`
}

type ArticleUpdateRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
	ImageID int    `json:"imageId" validate:"required"`
}

func NewArticleHandler(articleUsecase usecase.IfArticleUsecase) IfArticleHandler {
	return &ArticleHandler{
		articleUsecase,
	}
}

func (h *ArticleHandler) Create(c echo.Context) error {
	var req ArticleCreateRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := c.Validate(&req); err != nil {
		return handleError(domain.ErrBadRequest, err)
	}

	if err := h.articleUsecase.Create(
		c.Request().Context(),
		input.NewArticle(
			0,
			req.Title,
			req.Content,
			req.ImageID,
			c.Get("userID").(int),
		),
	); err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}

func (h *ArticleHandler) List(c echo.Context) error {
	es, err := h.articleUsecase.List(c.Request().Context(), c.Get("userID").(int))
	if err != nil {
		return err
	}
	return presenter.BuildArticleListResponse(c, es)
}

func (h *ArticleHandler) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("articleId"))
	if err != nil {
		return err
	}
	e, err := h.articleUsecase.Get(c.Request().Context(), id, c.Get("userID").(int))
	if err != nil {
		return err
	}
	return presenter.BuildArticleResponse(c, e)
}

func (h *ArticleHandler) Update(c echo.Context) error {
	var req ArticleUpdateRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := c.Validate(&req); err != nil {
		return handleError(domain.ErrBadRequest, err)
	}
	id, err := strconv.Atoi(c.Param("articleId"))
	if err != nil {
		return err
	}

	if err := h.articleUsecase.Update(
		c.Request().Context(),
		input.NewArticle(
			id,
			req.Title,
			req.Content,
			req.ImageID,
			c.Get("userID").(int),
		),
	); err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *ArticleHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("articleId"))
	if err != nil {
		return err
	}

	if err = h.articleUsecase.Delete(c.Request().Context(), id, c.Get("userID").(int)); err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
