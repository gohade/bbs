package qa

import (
	"bbs/app/http/middleware/auth"
	provider "bbs/app/provider/qa"
	"github.com/gohade/hade/framework/gin"
	"github.com/pkg/errors"
)

// QuestionCreate 代表创建问题
func (api *QAApi) QuestionCreate(c *gin.Context)  {
	qaService := c.MustMake(provider.QaKey).(provider.Service)
	type Param struct {
		Title string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}
	param := &Param{}
	if err := c.ShouldBind(param); err != nil {
		c.AbortWithError(404, err); return
	}

	user := auth.GetAuthUser(c)
	if user == nil {
		c.AbortWithError(500, errors.New("无权限操作")); return
	}

	question := &provider.Question{
		Title:     param.Title,
		Context:   param.Content,
		AuthorID:  user.ID,
	}
	if err := qaService.PostQuestion(c, question); err != nil {
		c.AbortWithError(500, err); return
	}

	c.ISetOkStatus().IText("操作成功")
}
