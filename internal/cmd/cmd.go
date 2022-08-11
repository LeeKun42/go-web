package cmd

import (
	"context"
	"fmt"
	"go-web/internal/controller"
	"go-web/internal/cron"
	"go-web/internal/job"
	"go-web/internal/service/user"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

/**
检查是否登录，检查不通过响应{"code":401, "message":"xxx"}
*/
func checkLogin(r *ghttp.Request) {
	user.Auth().MiddlewareFunc()(r)
	r.Middleware.Next()
}

func requestLog(r *ghttp.Request) {
	g.Log().Info(r.GetCtx(), r.GetUrl(), r.Header, r.GetBodyString())
	r.Middleware.Next()
}

var (
	Main = gcmd.Command{
		Name:        "main",
		Brief:       "start http server",
		Description: "start web server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			//定时任务
			cron.Schedule(ctx)

			//web服务
			s := g.Server()
			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareCORS, ghttp.MiddlewareHandlerResponse, requestLog)
				group.POST("/user/login", controller.UserController.Login)
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(checkLogin)
					group.Group("/user", func(group *ghttp.RouterGroup) {
						group.GET("/info", controller.UserController.GetLoginUserInfo)
						group.GET("/lists", controller.UserController.GetList)
						group.GET("/token/refresh", controller.UserController.RefreshToken)
						group.POST("/logout", controller.UserController.Logout)
					})
				})

			})
			s.Run()

			return nil
		},
	}

	Test = &gcmd.Command{
		Name:  "test:work",
		Brief: "test Brief",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			fmt.Println("this is command test")
			//任务投递
			job.NewSmsJob("13802990402", "1").Dispatch(ctx)
			job.NewSmsJob("13802990402", "2").Dispatch(ctx)
			job.NewSmsJob("13802990402", "3").Dispatch(ctx)
			job.NewSmsJob("13802990402", "4").Dispatch(ctx)
			job.NewSmsJob("13802990402", "5").Dispatch(ctx)
			job.NewEmailJob("980984232@qq.com", "1").Dispatch(ctx)
			job.NewEmailJob("980984232@qq.com", "2").Dispatch(ctx)
			job.NewEmailJob("980984232@qq.com", "3").Dispatch(ctx)
			job.NewEmailJob("980984232@qq.com", "4").Dispatch(ctx)
			job.NewEmailJob("980984232@qq.com", "5").Dispatch(ctx)
			return nil
		},
	}

	QueueWork = &gcmd.Command{
		Name:  "queue:work",
		Brief: "start queue work with --name",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			name := parser.GetOpt("name").String()
			job.QueueWork(ctx, name)
			return nil
		},
	}

	QueueRetry = &gcmd.Command{
		Name:  "queue:retry",
		Brief: "queue retry with fail job id",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			id := parser.GetOpt("id").Int()
			job.QueueRetry(ctx, id)
			return nil
		},
	}
)
