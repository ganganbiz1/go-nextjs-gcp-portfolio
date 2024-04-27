package entity

type User struct {
	ID                 int
	Email              string
	Name               string
	FirebaseUID        string
	FirebaseProviderID string
}

func NewUser(
	id int,
	email,
	name,
	firebaseUID,
	firebaseProviderID string,
) *User {
	return &User{
		ID:                 id,
		Email:              email,
		Name:               name,
		FirebaseUID:        firebaseUID,
		FirebaseProviderID: firebaseProviderID,
	}
}

func (e *User) SetFirebaseUID(uid string) {
	e.FirebaseUID = uid
}

func (e *User) SetFirebaseProviderID(providerID string) {
	e.FirebaseProviderID = providerID
}
