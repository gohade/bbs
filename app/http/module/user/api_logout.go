package user

import (
	"bbs/app/http/middleware/auth"
	provider "bbs/app/provider/user"
	"github.com/gohade/hade/framework/gin"
	"github.com/pkg/errors"
)

// Logout 代表登出
func (api *UserApi) Logout(c *gin.Context)  {
	authUser := auth.GetAuthUser(c)
	if authUser == nil {
		c.AbortWithError(500, errors.New("用户未登录")); return
	}

	userService := c.MustMake(provider.UserKey).(provider.Service)
	if err := userService.Logout(c, authUser); err != nil {
		c.AbortWithError(500, errors.New("用户未登录")); return
	}
	c.ISetOkStatus(); return
}

