package user

import (
	"bbs/app/http/middleware/auth"
	provider "bbs/app/provider/user"
	"github.com/gohade/hade/framework/gin"
	"github.com/pkg/errors"
)

// Logout 代表登出
// @Summary 用户登出
// @Description 调用表示用户登出
// @Accept  json
// @Produce  json
// @Tags user
// @Success 200 {string} Message "用户登出成功"
// @Router /user/logout [get]
func (api *UserApi) Logout(c *gin.Context)  {
	authUser := auth.GetAuthUser(c)
	if authUser == nil {
		c.AbortWithError(500, errors.New("用户未登录")); return
	}

	userService := c.MustMake(provider.UserKey).(provider.Service)
	if err := userService.Logout(c, authUser); err != nil {
		c.AbortWithError(500, errors.New("用户未登录")); return
	}
	c.ISetOkStatus().IText("用户登出成功"); return
}

