package user

import (
	provider "bbs/app/provider/user"
	"github.com/gohade/hade/framework/gin"
	"github.com/pkg/errors"
)

// Verify 代表验证注册信息
func (api *UserApi) Verify(c *gin.Context)  {
	// 验证参数
	userService := c.MustMake(provider.UserKey).(provider.Service)
	type Param struct {
		Token string `json:"token" binding:"required"`
	}
	param := &Param{}
	if err := c.ShouldBind(param); err != nil {
		c.AbortWithError(404, err); return
	}

	verified, err := userService.VerifyRegister(c, param.Token)
	if err != nil {
		c.AbortWithError(500, err); return
	}

	if !verified {
		c.AbortWithError(500, errors.New("验证错误")); return
	}

	// 输出
	c.ISetOkStatus().IText("请重新登录")
}
