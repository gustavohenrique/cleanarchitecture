package uuid

import "github.com/google/uuid"

func NewV4() string {
	u, err := uuid.NewRandom()
	if err != nil {
		return ""
	}
	return u.String()
}

func IsInvalid(s string) bool {
	_, err := uuid.Parse(s)
	return err != nil
}
