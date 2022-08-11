package cron

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/os/gcron"
)

func Schedule(ctx context.Context) {
	_, errCron := gcron.Add(ctx, "*/5 * * * * *", func(ctx context.Context) {
		fmt.Println(gtime.Datetime() + " Every five second")
	}, "MySecondCronJob")
	if errCron != nil {
		panic(errCron)
	}
}
