package user

import (
	provider "bbs/app/provider/user"
	"github.com/gohade/hade/framework/gin"
	"github.com/pkg/errors"
	"time"
)

// Register 注册
func (api *UserApi) Register(c *gin.Context)  {
	// 验证参数
	userService := c.MustMake(provider.UserKey).(provider.Service)
	type Param struct {
		UserName string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required,gte=6"`
		Email string `json:"email" binding:"required,gte=6"`
	}
	param := &Param{}
	if err := c.ShouldBind(param); err != nil {
		c.AbortWithError(404, err); return
	}

	// 登录
	model := &provider.User{
		UserName:  param.UserName,
		Password:  param.Password,
		Email: param.Email,
		CreatedAt: time.Now(),
	}
	// 注册
	userWithToken, err := userService.Register(c, model)
	if err != nil {
		c.AbortWithError(500, err); return
	}
	if userWithToken == nil {
		c.AbortWithError(500, errors.New("注册失败")); return
	}

	if err := userService.SendRegisterMail(c, userWithToken); err != nil {
		c.AbortWithError(500, errors.New("发送电子邮件失败")); return
	}

	c.ISetOkStatus().IText("操作成功")
	// 输出
	c.IJson(map[string]interface{}{"token": userWithToken.Token}); return
}

