package cryptus_test

import (
	"testing"

	"{{ .ProjectName }}/src/shared/cryptus"
	"{{ .ProjectName }}/src/shared/test/assert"
)

func TestRsaGenerateKeyPairAndEncryptAndDecrypt(t *testing.T) {
	priv, pub, err := cryptus.GenerateRsaKeyPair(2048)
	assert.Nil(t, err)

	password := "mypassword"
	encrypted, err := cryptus.EncryptRsa([]byte(password), pub)
	assert.Nil(t, err)

	decrypted, err := cryptus.DecryptRsa(encrypted, priv)
	assert.Nil(t, err)

	assert.Equal(t, string(decrypted), password)
}

func TestPbkdf2(t *testing.T) {
	password := "mypassword"
	hash, salt := cryptus.HashPbkdf2(password, []byte(""))
	newHash, salt := cryptus.HashPbkdf2(password, salt)
	assert.Equal(t, string(hash), string(newHash))
}
