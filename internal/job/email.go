package job

import (
	"context"
	"fmt"
)

type Email struct {
	Address string `json:"address"`
	Content string `json:"content"`
}

var email Email
var EmailQueueName = "email"

func NewEmailJob(address string, content string) *Email {
	email = Email{Address: address, Content: content}
	return &email
}

func (e *Email) Dispatch(ctx context.Context) {
	Dispatch(ctx, EmailQueueName, e)
}

func (e *Email) Handle(ctx context.Context) {
	fmt.Printf("email job handle: address is %s, content is %s\n", e.Address, e.Content)
}
