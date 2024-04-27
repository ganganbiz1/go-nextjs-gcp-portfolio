package input

import "github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/entity"

type Article struct {
	ID      int
	Title   string
	Content string
	ImageID int
	UserID  int
}

func NewArticle(
	id int,
	title,
	content string,
	imageID,
	userID int,
) *Article {
	return &Article{
		ID:      id,
		Title:   title,
		Content: content,
		ImageID: imageID,
		UserID:  userID,
	}
}

func (d *Article) ToEntity() *entity.Article {
	return entity.NewArticle(
		d.ID,
		d.Title,
		d.Content,
		d.ImageID,
		d.UserID,
	)
}
