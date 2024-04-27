package model

import (
	"time"

	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/entity"
	"gorm.io/gorm"
)

type User struct {
	ID                 int
	Email              string
	Name               string
	FirebaseUID        string
	FirebaseProviderID string
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          *gorm.DeletedAt
}

type Users []*User

func (m *User) ToEntity() *entity.User {
	return entity.NewUser(
		m.ID,
		m.Email,
		m.Name,
		m.FirebaseUID,
		m.FirebaseProviderID,
	)
}

func (ms Users) ToEntity() []*entity.User {
	es := make([]*entity.User, 0, len(ms))
	for _, m := range ms {
		es = append(es, m.ToEntity())
	}

	return es
}

func (m *User) TableName() string {
	return "users"
}
