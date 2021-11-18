package user

import (
	"bbs/app/http/middleware/auth"
	"bbs/app/provider/user"
	"github.com/gohade/hade/framework/gin"
)

type UserApi struct {}

// RegisterRoutes 注册路由
func RegisterRoutes(r *gin.Engine) error {
	api := &UserApi{}
	if !r.IsBind(user.UserKey) {
		r.Bind(&user.UserProvider{})
	}

	r.POST("/user/login", api.Login)
	r.GET("/user/logout", auth.AuthMiddleware(), api.Logout)

	return nil
}

