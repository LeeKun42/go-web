package model

import "go-web/internal/model/entity"

type UserAuthInput struct {
	Mobile   string
	Password string
}

type GetUserInput struct {
	Name   string
	Mobile string
	PaginateInput
}

type GetUserOutput struct {
	Users []*entity.User `json:"users"      `
	PaginateOutput
}
