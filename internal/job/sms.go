package job

import (
	"context"
	"fmt"
)

type Sms struct {
	Mobile  string `json:"mobile"`
	Content string `json:"content"`
}

var sms Sms
var SmsQueueName = "sms"

func NewSmsJob(mobile string, content string) *Sms {
	sms = Sms{Mobile: mobile, Content: content}
	return &sms
}

func (s *Sms) Dispatch(ctx context.Context) {
	Dispatch(ctx, SmsQueueName, s)
}

func (s *Sms) Handle(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil { //执行出错
			//失败队列处理
		}
	}()
	fmt.Printf("sms job handle: mobile is %s, content is %s\n", s.Mobile, s.Content)
}
