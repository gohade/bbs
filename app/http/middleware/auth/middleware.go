package auth

import (
	"bbs/app/provider/user"
	"github.com/gohade/hade/framework/gin"
)

// AuthMiddleware 登录中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil || token == "" {
			c.ISetStatus(401).IText("请登录后操作"); return
		}

		userService := c.MustMake(user.UserKey).(user.Service)
		user, err := userService.VerifyLogin(c, token)
		if err != nil || user == nil {
			c.ISetStatus(401).IText("请登录后操作"); return
		}

		c.Set("auth_user", user)

		c.Next()
	}
}

// GetAuthUser 获取已经验证的用户
func GetAuthUser(c *gin.Context) *user.User {
	t, exist:= c.Get("auth_user")
	if !exist {
		return nil
	}
	return t.(*user.User)
}

