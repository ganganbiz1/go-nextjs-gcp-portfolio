package presenter

import (
	"net/http"

	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/entity"
	"github.com/labstack/echo/v4"
)

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	ImageID int    `json:"imageId"`
	UserID  int    `json:"userId"`
}

func BuildArticleResponse(c echo.Context, e *entity.Article) error {
	return BuildSuccessResponse(c, http.StatusOK, &Article{
		ID:      e.ID,
		Title:   e.Title,
		Content: e.Content,
		ImageID: e.ImageID,
		UserID:  e.UserID,
	})
}

func BuildArticleListResponse(c echo.Context, es []*entity.Article) error {
	list := make([]*Article, 0, len(es))

	for _, e := range es {
		list = append(list, &Article{
			ID:      e.ID,
			Title:   e.Title,
			Content: e.Content,
			ImageID: e.ImageID,
			UserID:  e.UserID,
		})
	}

	return BuildSuccessResponse(c, http.StatusOK, list)
}
