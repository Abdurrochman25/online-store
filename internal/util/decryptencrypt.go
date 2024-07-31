package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var ErrDecryptDataToShort = errors.New("decrypt: plaint text to short")

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hashedPassword), err
}

func CheckPasswordHash(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func DecryptData(plainText, key string) string {
	if plainText == "" {
		return plainText
	}
	text, errDecodeString := base64.StdEncoding.DecodeString(plainText)
	if errDecodeString != nil {
		return errDecodeString.Error()
	}
	block, errNewCipher := aes.NewCipher([]byte(key))
	if errNewCipher != nil {
		return errNewCipher.Error()
	}
	if len(text) < aes.BlockSize {
		return ErrDecryptDataToShort.Error()
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)

	return string(text)
}
