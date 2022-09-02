package cryptus

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/crypto/pbkdf2"

	"{{ .ProjectName }}/src/shared/uuid"
)

const (
	HASH_AND_SALT_SEPARATOR = "$"
	SALT_SIZE               = 32
)

func RandomSecretKey() string {
	nonce := randomBase64Nonce()
	key := strings.ReplaceAll(uuid.NewV4(), "-", "") // Key should have 32 bytes
	return nonce + key
}

func HashIt(password string, salt []byte) string {
	hash := pbkdf2.Key([]byte(password), salt, 4096, 64, sha256.New)
	return fmt.Sprintf("%x%s%x", hash, HASH_AND_SALT_SEPARATOR, salt)
}

func ComparePasswordAndHash(password, hash string) (match bool, err error) {
	values := strings.Split(hash, HASH_AND_SALT_SEPARATOR)
	if len(values) != 2 {
		return false, fmt.Errorf("hash doesnt contains salt")
	}
	salt, _ := hex.DecodeString(values[1])
	newValues := strings.Split(HashIt(password, salt), HASH_AND_SALT_SEPARATOR)
	oldHash := values[0]
	newHash := newValues[0]
	if newHash == oldHash {
		return true, nil
	}
	return false, fmt.Errorf(newHash)
}

func Salt() []byte {
	b := make([]byte, SALT_SIZE)
	rand.Read(b)
	return b
}

func ToBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func FromBase64(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}

func IsBase64(str string) bool {
	_, err := FromBase64(str)
	return err == nil
}

func randomBase64Nonce() string {
	rand.Seed(time.Now().UnixNano())
	nonce := make([]byte, 12) // AES nonce 96 bits = 12 bytes = 16 bytes in base64
	_, err := rand.Read(nonce)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(nonce)
}
