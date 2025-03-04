package hasherpkg

import (
	"crypto/sha256"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type HasherInterface interface {
	Hash(password string) string
	Check(password string, hash string) bool
}

type CryptHasher struct {
}

func NewCryptHasher() CryptHasher {
	return CryptHasher{}
}

func (b *CryptHasher) Hash(password string) string {
	hashedpassword := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", hashedpassword)
}

func (b *CryptHasher) Check(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Fatal(err)
	}
	return true
}
