package user

import (
	provider "bbs/app/provider/user"
	"github.com/gohade/hade/framework/gin"
	"github.com/pkg/errors"
	"time"
)

type registerParam struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,gte=6"`
	Email string `json:"email" binding:"required,gte=6"`
}

// Register 注册
// @Summary 用户注册
// @Description 用户注册接口
// @Accept  json
// @Produce  json
// @Tags user
// @Param registerParam body registerParam true "注册参数"
// @Success 200 {string} Message "注册成功"
// @Router /user/register [post]
func (api *UserApi) Register(c *gin.Context)  {
	// 验证参数
	userService := c.MustMake(provider.UserKey).(provider.Service)

	param := &registerParam{}
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

	c.ISetOkStatus().IText("注册成功，请前往邮箱查看邮件"); return
}

