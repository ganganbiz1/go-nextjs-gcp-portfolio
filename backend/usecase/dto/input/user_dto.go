package input

import "github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/entity"

type User struct {
	ID                 int
	Email              string
	Name               string
	Password           string
	IDToken            string
	FirebaseUID        string
	FirebaseProviderID string
}

func NewUser(
	id int,
	email,
	name,
	password,
	idToken,
	firebaseUID,
	firebaseProviderID string,
) *User {
	return &User{
		ID:                 id,
		Email:              email,
		Name:               name,
		Password:           password,
		IDToken:            idToken,
		FirebaseUID:        firebaseUID,
		FirebaseProviderID: firebaseProviderID,
	}
}

func (d *User) ToEntity() *entity.User {
	return entity.NewUser(
		d.ID,
		d.Email,
		d.Name,
		d.FirebaseUID,
		d.FirebaseProviderID,
	)
}
