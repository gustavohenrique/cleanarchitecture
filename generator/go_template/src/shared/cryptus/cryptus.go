package cryptus

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

var hash = sha256.New()

func FromBytesToBase64(val []byte) string {
	return base64.StdEncoding.EncodeToString(val)
}

func ToBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func FromBase64(str string) ([]byte, error) {
	if strings.TrimSpace(str) == "" {
		return []byte(""), fmt.Errorf("empty base64 string")
	}
	return base64.StdEncoding.DecodeString(str)
}

func IsBase64(str string) bool {
	_, err := FromBase64(str)
	return err == nil
}

func SignRsa(plaintext, privkey string) (string, error) {
	privateKey, err := ExtractPrivateKeyFromBase64(privkey)
	if err != nil {
		return "", err
	}
	msgHashSum := hashSum(plaintext)
	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, msgHashSum, nil)
	if err != nil {
		return "", err
	}
	return ToBase64(string(signature)), nil
}

func ExtractPrivateKeyFromBase64(privkey string) (*rsa.PrivateKey, error) {
	privkeyBytes, err := FromBase64(privkey)
	if err != nil {
		return nil, err
	}
	return getRsaPrivateKeyFrom(privkeyBytes)
}

func ExtractPublicKeyFromBase64(pubkey string) (*rsa.PublicKey, error) {
	pubkeyBytes, err := FromBase64(pubkey)
	if err != nil {
		return nil, err
	}
	return getRsaPublicKeyFrom(pubkeyBytes)
}

func VerifyRsa(plaintext, pubkey, sig string) error {
	signature, err := FromBase64(sig)
	if err != nil {
		return err
	}
	pubkeyBytes, err := FromBase64(pubkey)
	if err != nil {
		return err
	}
	publicKey, err := getRsaPublicKeyFrom(pubkeyBytes)
	if err != nil {
		return err
	}
	msgHashSum := hashSum(plaintext)
	return rsa.VerifyPSS(publicKey, crypto.SHA256, msgHashSum, signature, nil)
}

func EncryptRsa(ciphertext []byte, pub []byte) ([]byte, error) {
	publicKey, err := getRsaPublicKeyFrom(pub)
	if err != nil {
		return []byte(""), err
	}
	encrypted, err := rsa.EncryptOAEP(hash, rand.Reader, publicKey, ciphertext, nil)
	return encrypted, err
}

func DecryptRsa(ciphertext []byte, priv []byte) ([]byte, error) {
	privateKey, err := getRsaPrivateKeyFrom(priv)
	if err != nil {
		return []byte(""), err
	}
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, ciphertext, nil)
	return plaintext, err
}

func GenerateRsaKeyPair(bitSize int) ([]byte, []byte, error) {
	var empty []byte
	key, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return empty, empty, err
	}

	privBytes, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		return empty, empty, err
	}
	privPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privBytes,
		},
	)
	pub := key.Public()
	pubBytes, err := x509.MarshalPKIXPublicKey(pub.(*rsa.PublicKey))
	if err != nil {
		return empty, empty, err
	}
	pubPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubBytes,
		},
	)
	return privPEM, pubPEM, nil
}

func ToHex(val []byte) string {
	return hex.EncodeToString(val)
}

func HashPbkdf2(plaintext string, salt []byte) ([]byte, []byte) {
	saltSize := 32
	if len(salt) != saltSize {
		salt = make([]byte, saltSize)
		io.ReadFull(rand.Reader, salt)
	}
	iterations := 100
	dk := pbkdf2.Key([]byte(plaintext), salt, iterations, saltSize, sha256.New)
	return dk, salt
}

func getRsaPrivateKeyFrom(priv []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(priv)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return nil, err
		}
	}
	privkey, err := x509.ParsePKCS8PrivateKey(b)
	if err != nil {
		return nil, err
	}
	return privkey.(*rsa.PrivateKey), nil
}

func getRsaPublicKeyFrom(pub []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pub)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return nil, err
		}
	}
	pubkey, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		return nil, err
	}
	return pubkey.(*rsa.PublicKey), nil
}

func hashSum(plaintext string) []byte {
	algo := sha256.New()
	_, err := algo.Write([]byte(plaintext))
	if err != nil {
		return []byte("")
	}
	msgHashSum := algo.Sum(nil)
	return msgHashSum
}
