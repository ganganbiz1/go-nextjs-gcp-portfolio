package model

import (
	"time"

	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/entity"
	"gorm.io/gorm"
)

type Article struct {
	ID        int
	Title     string
	Content   string
	ImageID   int
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *gorm.DeletedAt
}

type Articles []*Article

func (m *Article) ToEntity() *entity.Article {
	return entity.NewArticle(
		m.ID,
		m.Title,
		m.Content,
		m.ImageID,
		m.UserID,
	)
}

func (ms Articles) ToEntity() []*entity.Article {
	es := make([]*entity.Article, 0, len(ms))
	for _, m := range ms {
		es = append(es, m.ToEntity())
	}

	return es
}

func (m *Article) TableName() string {
	return "articles"
}
