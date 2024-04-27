package entity

type TokenInfo struct {
	AuthTime int
	Issuer   string
	Audience string
	Expires  int
	IssuedAt int
	Subject  string
	UID      string
}

func NewTokenInfo(
	authTime int,
	issuer,
	audience string,
	expires,
	issuedAt int,
	subject,
	uid string,
) *TokenInfo {
	return &TokenInfo{
		AuthTime: authTime,
		Issuer:   issuer,
		Audience: audience,
		Expires:  expires,
		IssuedAt: issuedAt,
		Subject:  subject,
		UID:      uid,
	}
}
