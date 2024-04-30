package entity

type User struct {
	ID                 int
	Email              string
	Name               string
	FirebaseUID        string
	FirebaseProviderID string

	Articles []*Article
}

func NewUser(
	id int,
	email,
	name,
	firebaseUID,
	firebaseProviderID string,
	articles []*Article,
) *User {
	return &User{
		ID:                 id,
		Email:              email,
		Name:               name,
		FirebaseUID:        firebaseUID,
		FirebaseProviderID: firebaseProviderID,

		Articles: articles,
	}
}

func (e *User) SetFirebaseUID(uid string) {
	e.FirebaseUID = uid
}

func (e *User) SetFirebaseProviderID(providerID string) {
	e.FirebaseProviderID = providerID
}
