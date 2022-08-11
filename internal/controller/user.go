package controller

import (
	"fmt"
	"go-web/internal/model"
	"go-web/internal/service/user"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf-jwt/v2/example/api"
	"github.com/gogf/gf/v2/net/ghttp"
)

type userController struct{}

var UserController = userController{}

func (uc *userController) GetList(request *ghttp.Request) {
	var input model.GetUserInput
	if err := request.Parse(&input); err != nil {
		request.Response.WriteJsonExit(err)
	}
	uid := gconv.Int(user.Auth().GetIdentity(request.Context()))
	fmt.Println(uid)
	ret := user.UserService.GetList(request.Context(), input)
	Response(request).success(ret)
}

func (uc *userController) GetLoginUserInfo(request *ghttp.Request) {
	uid := gconv.Int(user.Auth().GetIdentity(request.Context()))
	ret := user.UserService.GetUserInfoByUid(request.Context(), uid)
	config, _ := g.Cfg().Get(request.Context(), "redis")
	fmt.Println(config)
	Response(request).success(ret)
}

func (uc *userController) Login(request *ghttp.Request) {
	res := &api.AuthLoginRes{}
	res.Token, res.Expire = user.Auth().LoginHandler(request.Context())
	request.Response.WriteJsonExit(res)
}

func (uc *userController) RefreshToken(request *ghttp.Request) {
	res := &api.AuthRefreshTokenRes{}
	res.Token, res.Expire = user.Auth().RefreshHandler(request.Context())
	request.Response.WriteJsonExit(res)
}

func (uc *userController) Logout(request *ghttp.Request) {
	user.Auth().LogoutHandler(request.Context())
	Response(request).success(nil)
}
