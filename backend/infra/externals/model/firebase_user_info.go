package model

import "github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/entity"

type FirebaseUserInfo struct {
	Email      string
	UID        string
	ProviderID string
}

func NewFirebaseUserInfo(email, uid, providerID string) *FirebaseUserInfo {
	return &FirebaseUserInfo{
		Email:      email,
		UID:        uid,
		ProviderID: providerID,
	}
}

func (m *FirebaseUserInfo) ToUserEntity() *entity.User {
	return entity.NewUser(0, m.Email, "", m.UID, m.ProviderID)
}
