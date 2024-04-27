package entity

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
