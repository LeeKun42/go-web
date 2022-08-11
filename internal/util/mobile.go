package util

import (
	"go-web/internal/consts"

	"github.com/marspere/goencrypt"
)

var Mobile = mobile{}

type mobile struct{}

func (m *mobile) Encrypt(mobile string) string {
	key := []byte(consts.MOBILE_SALT)
	ivByte := [16]byte{}
	aesCipher, err := goencrypt.NewAESCipher(key, ivByte[:], goencrypt.ECBMode, goencrypt.Pkcs7, goencrypt.PrintBase64)
	if err != nil {
		panic(err)
	}
	mobileEncrypt, err := aesCipher.AESEncrypt([]byte(mobile))
	if err != nil {
		panic(err)
	}
	return mobileEncrypt
}

func (m *mobile) Decrypt(mobileEncrypt string) string {
	key := []byte(consts.MOBILE_SALT)
	ivByte := [16]byte{}
	aesCipher, err := goencrypt.NewAESCipher(key, ivByte[:], goencrypt.ECBMode, goencrypt.Pkcs7, goencrypt.PrintBase64)
	if err != nil {
		panic(err)
	}
	mobile, err := aesCipher.AESDecrypt(mobileEncrypt)
	if err != nil {
		panic(err)
	}
	return mobile
}
