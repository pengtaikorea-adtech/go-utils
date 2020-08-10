package secure

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"io"

	"github.com/google/uuid"
)

// GenerateKey generate Random Key
func GenerateKey() string {
	// build uuid
	var salt string
	if uid, err := uuid.NewUUID(); err == nil {
		salt = uid.String()
	} else {
		salt = uuid.Nil.String()
	}
	return Sha256(salt)
}

// GenerateRandom512 returns sha512 hashed random string
func GenerateRandom512() string {
	nonce := make([]byte, 512)
	io.ReadFull(rand.Reader, nonce)
	return Sha512(string(nonce))
}

// Sha512 - convert sha512 hash string
func Sha512(token string) string {
	bytes := []byte(token)
	hashed := sha512.Sum512(bytes)
	return hex.EncodeToString(hashed[:])
}

// Sha256 - convert sha256 hash string
func Sha256(token string) string {
	bytes := []byte(token)
	hashed := sha256.Sum256(bytes)
	return hex.EncodeToString(hashed[:])
}

// AESEncrypt the string, AES256 by key - base64 encoded
func AESEncrypt(key string, payload string) string {
	keyDecode, _ := hex.DecodeString(key)
	block, _ := aes.NewCipher(keyDecode)
	// build iv
	iv := make([]byte, 12)
	io.ReadFull(rand.Reader, iv)

	encrypter, _ := cipher.NewGCM(block)
	plaintext := []byte(payload)
	crypt := encrypter.Seal(nil, iv, plaintext, nil)
	outs := make([]byte, len(iv)+len(crypt))

	copy(outs, iv)
	copy(outs[len(iv):], crypt)

	return base64.RawURLEncoding.EncodeToString(outs)
}

// AESDecrypt the string, base64 encoder
func AESDecrypt(key string, crypt string) string {
	keyDecode, _ := hex.DecodeString(key)
	block, _ := aes.NewCipher(keyDecode)

	bytes, _ := base64.RawURLEncoding.DecodeString(crypt)

	iv := bytes[0:12]
	payload := bytes[12:]

	decrypter, _ := cipher.NewGCM(block)
	plains, _ := decrypter.Open(nil, iv, payload, nil)

	return string(plains)
}
