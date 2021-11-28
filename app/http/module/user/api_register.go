package user

import (
	provider "bbs/app/provider/user"
	"fmt"
	"github.com/gohade/hade/framework/contract"
	"github.com/gohade/hade/framework/gin"
	"time"
)

type registerParam struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,gte=6"`
	Email    string `json:"email" binding:"required,gte=6"`
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
func (api *UserApi) Register(c *gin.Context) {
	// 验证参数
	userService := c.MustMake(provider.UserKey).(provider.Service)
	logger := c.MustMake(contract.LogKey).(contract.Log)

	param := &registerParam{}
	if err := c.ShouldBind(param); err != nil {
		c.ISetStatus(400).IText("参数错误")
		return
	}

	// 登录
	model := &provider.User{
		UserName:  param.UserName,
		Password:  param.Password,
		Email:     param.Email,
		CreatedAt: time.Now(),
	}
	// 注册
	userWithToken, err := userService.Register(c, model)
	if err != nil {
		logger.Error(c, err.Error(), map[string]interface{}{
			"stack": fmt.Sprintf("%+v", err),
		})
		c.ISetStatus(500).IText(err.Error())
		return
	}
	if userWithToken == nil {
		c.ISetStatus(500).IText("注册失败")
		return
	}

	if err := userService.SendRegisterMail(c, userWithToken); err != nil {

		c.ISetStatus(500).IText("发送电子邮件失败")
		return
	}

	c.ISetOkStatus().IText("注册成功，请前往邮箱查看邮件")
	return
}
