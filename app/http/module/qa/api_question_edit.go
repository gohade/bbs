package qa

import (
	"bbs/app/http/middleware/auth"
	provider "bbs/app/provider/qa"
	"github.com/gohade/hade/framework/gin"
	"github.com/pkg/errors"
)

// QuestionEdit 代表编辑问题
func (api *QAApi) QuestionEdit(c *gin.Context)  {
	qaService := c.MustMake(provider.QaKey).(provider.Service)
	type Param struct {
		ID int64 `json:"id" binding:"required"`
		Title string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}
	param := &Param{}
	if err := c.ShouldBind(param); err != nil {
		c.AbortWithError(404, err); return
	}

	questionOld, err := qaService.GetQuestion(c, param.ID)
	if err != nil || questionOld == nil {
		c.AbortWithError(500, errors.New("操作的问题不存在"))
	}

	user := auth.GetAuthUser(c)
	if user == nil  || user.ID != questionOld.AuthorID {
		c.AbortWithError(500, errors.New("无权限操作")); return
	}

	question := &provider.Question{
		ID: param.ID,
		Title:     param.Title,
		Context:   param.Content,
	}
	if err := qaService.UpdateQuestion(c, question); err != nil {
		c.AbortWithError(500, err); return
	}

	c.ISetOkStatus().IText("操作成功")
}

