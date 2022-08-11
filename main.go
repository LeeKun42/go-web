package main

import (
	_ "go-web/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"go-web/internal/cmd"
)

func main() {
	ctx := gctx.New()

	//添加子命令
	cmd.Main.AddCommand(cmd.Test, cmd.QueueWork, cmd.QueueRetry)

	//运行主服务
	cmd.Main.Run(ctx)
}
