package job

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

type QueueJob interface {
	Handle(ctx context.Context)
}

var queuePrefix = "queue"

// QueueMap 所有队列map：队列名称=》队列运行器
var QueueMap = map[string]QueueJob{
	SmsQueueName:   &Sms{Mobile: "", Content: ""},
	EmailQueueName: &Email{Address: "", Content: ""},
}

func Dispatch(ctx context.Context, name string, job QueueJob) {
	json := gjson.New(job)
	str, _ := json.ToJsonString()
	key := queuePrefix + ":" + name
	g.Redis("db").Do(ctx, "RPUSH", key, str)
	fmt.Println("sms job dispatch", str)
}

// QueueWork 运行队列消费端
func QueueWork(ctx context.Context, name string) {
	conn, err := g.Redis("db").Conn(ctx)
	if err != nil {
		panic(err)
	}
	smsQueueKey := "queue:" + name
	for {
		fmt.Println("queue " + name + " work pop job")
		result, err := conn.Do(ctx, "BLPOP", smsQueueKey, 5)
		if err != nil { //获取出错
			time.Sleep(time.Second * 1)
			fmt.Println("err continue", err)
			continue
		}
		if result.IsNil() { //结果为空
			time.Sleep(time.Second * 1)
			fmt.Println("blpop result is empty continue")
			continue
		}
		fmt.Println("queue " + name + " work start")
		//redis获取的数据转换为数组，并获取第二个值。（注：BLPOP命令返回值比较特殊第一个返回值为list的key，第二个返回为lpop出来的值）
		json, _ := gjson.DecodeToJson(result.Array()[1])
		job := QueueMap[name]
		json.Scan(job)
		//执行具体业务逻辑处理（可用协程）
		go job.Handle(ctx)
		fmt.Println("queue " + name + " work complete")
	}
}

// QueueRetry 重试错误消费队列
func QueueRetry(ctx context.Context, jobId int) {

}
