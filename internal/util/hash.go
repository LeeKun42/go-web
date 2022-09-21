package util

import "golang.org/x/crypto/bcrypt"

var Hash = hash{}

type hash struct{}

// Make 对字符串进行hash/**
func (p *hash) Make(text string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hashed)
}

func (p *hash) Check(plainText string, hashedText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedText), []byte(plainText))
	if err != nil {
		return false
	} else {
		return true
	}
}

func (p *hash) NeedHash(hashedText string) bool {
	hasCost, err := bcrypt.Cost([]byte(hashedText))
	return err != nil || hasCost != bcrypt.DefaultCost
}
