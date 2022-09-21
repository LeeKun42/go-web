package user

import (
	"context"
	"go-web/internal/dao"
	"go-web/internal/model"
	"go-web/internal/model/do"
	"go-web/internal/model/entity"
	"go-web/internal/util"
	"time"

	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type userService struct{}

var UserService = userService{}

func (us *userService) Register(ctx context.Context, input do.User) (user *entity.User) {
	model := dao.User.Ctx(ctx)
	userId, _ := model.InsertAndGetId(input)
	model.WherePri(userId).Scan(&user)
	return user
}

func (us *userService) GetList(ctx context.Context, input model.GetUserInput) (res model.GetUserOutput) {
	model := dao.User.Ctx(ctx)

	if input.Name != "" {
		model = model.WhereLike(dao.User.Columns().Name, "%"+input.Name+"%")
	}

	if input.Mobile != "" {
		model = model.WhereLike(dao.User.Columns().Mobile, "%"+input.Mobile+"%")
	}
	total, _ := model.Count() //获取数据总量
	model = model.OrderDesc("id")
	model = model.Offset((input.CurrentPage - 1) * input.PageSize).Limit(input.PageSize)
	model.Scan(&res.Users)

	//分页信息计算
	res.Total = total
	res.CurrentPage = input.CurrentPage
	res.PageSize = input.PageSize
	res.PageCount = res.Total / res.PageSize
	if res.Total%res.PageSize > 0 {
		res.PageCount++
	}
	return res
}

func (us *userService) GetUserInfoByUid(ctx context.Context, uid int) (res *entity.User) {
	model := dao.User.Ctx(ctx)
	model.WherePri(uid).Scan(&res)
	return
}

// Jwt用户授权相关
var authService *jwt.GfJWTMiddleware

func Auth() *jwt.GfJWTMiddleware {
	return authService
}

func init() {
	auth := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte("3MoHngdbNkH9OtUYsqIsMuFMfoihH6qb4AnH5Ys1ozn2BNyfuxYWQYauvrU6sh2m"), //密钥
		Timeout:         time.Minute * 5000,                                                         //过期时间
		MaxRefresh:      time.Minute * 5000,                                                         //最长续期时间
		IdentityKey:     "id",
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Authenticator:   Authenticator,
		Unauthorized:    Unauthorized,
		PayloadFunc:     PayloadFunc,
		IdentityHandler: IdentityHandler,
	})
	authService = auth
}

// PayloadFunc is a callback function that will be called during login.
// Using this function it is possible to add additional payload data to the webtoken.
// The data is then made available during requests via c.Get("JWT_PAYLOAD").
// Note that the payload is not encrypted.
// The attributes mentioned on jwt.io can't be used as keys for the map.
// Optional, by default no additional data will be set.
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

// IdentityHandler get the identity from JWT and set the identity for every request
// Using this function, by r.GetParam("id") get identity
func IdentityHandler(ctx context.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return claims[authService.IdentityKey]
}

// Unauthorized is used to define customized Unauthorized callback function.
func Unauthorized(ctx context.Context, code int, message string) {
	r := g.RequestFromCtx(ctx)
	r.Response.WriteJson(g.Map{
		"code":    code,
		"message": message,
	})
	r.ExitAll()
}

// Authenticator is used to validate login parameters.
// It must return user data as user identifier, it will be stored in Claim Array.
// if your identityKey is 'id', your user data must have 'id'
// Check error (e) to determine the appropriate error message.
func Authenticator(ctx context.Context) (interface{}, error) {
	var (
		r         = g.RequestFromCtx(ctx)
		in        model.UserAuthInput
		loginUser *entity.User
	)
	if err := r.Parse(&in); err != nil {
		return "", err
	}
	userModel := dao.User.Ctx(ctx)

	err := userModel.Where(dao.User.Columns().Mobile, in.Mobile).Scan(&loginUser)

	if err != nil {
		panic(err)
	}
	if util.Hash.Check(in.Password, loginUser.Passwd) {
		return g.Map{"id": loginUser.Id, "name": loginUser.Name}, nil
	} else {
		return nil, gerror.New("密码错误")
	}
}
