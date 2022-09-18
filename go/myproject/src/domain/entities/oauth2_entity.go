package entities

import (
	"{{ .ProjectName }}/src/shared/cryptus"
	"{{ .ProjectName }}/src/shared/uuid"
)

type Oauth2 struct {
	ClientID     string
	ClientSecret string
	Token        string
	Hash         string
}

func NewOauth2(id string) Oauth2 {
	out := Oauth2{ClientID: id}
	if len(id) != 36 {
		out.ClientID = uuid.NewV4()
	}
	out.ClientSecret = cryptus.RandomSecretKey()
	out.Hash = cryptus.HashIt(out.ClientSecret, cryptus.Salt())
	return out
}

func (s Oauth2) ComparePasswordAndHash() (bool, error) {
	return cryptus.ComparePasswordAndHash(s.ClientSecret, s.Hash)
}
