package user

import (
	provider "bbs/app/provider/user"
	"github.com/gohade/hade/framework/gin"
	"github.com/pkg/errors"
)

type loginParam struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,gte=6"`
}

// Login 代表登录
// @Summary 用户登录
// @Description 用户登录接口
// @Accept  json
// @Produce  json
// @Tags user
// @Param loginParam body loginParam  true "login with param"
// @Success 200 {string} Token "token"
// @Router /user/login [post]
func (api *UserApi) Login(c *gin.Context)  {
	// 验证参数
	userService := c.MustMake(provider.UserKey).(provider.Service)

	param := &loginParam{}
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
	c.ISetOkStatus().IText(userWithToken.Token); return
}

