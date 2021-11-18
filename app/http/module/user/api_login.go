package user

import (
	provider "bbs/app/provider/user"
	"github.com/gohade/hade/framework/gin"
	"github.com/pkg/errors"
)

// Login 代表登录
func (api *UserApi) Login(c *gin.Context)  {
	// 验证参数
	userService := c.MustMake(provider.UserKey).(provider.Service)
	type Param struct {
		UserName string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required,gte=6"`
	}
	param := &Param{}
	if err := c.ShouldBind(param); err != nil {
		c.AbortWithError(404, err); return
	}

	// 登录
	model := &provider.User{
		UserName:  param.UserName,
		Password:  param.Password,
	}
	userWithToken, err := userService.Login(c, model)
	if err != nil {
		c.AbortWithError(500, err); return
	}
	if userWithToken == nil {
		c.AbortWithError(500, errors.New("用户不存在")); return
	}

	// 输出
	c.IJson(map[string]interface{}{"token": userWithToken.Token}); return
}

