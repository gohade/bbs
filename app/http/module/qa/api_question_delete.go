package qa

import (
	"bbs/app/http/middleware/auth"
	provider "bbs/app/provider/qa"
	"github.com/gohade/hade/framework/gin"
	"github.com/pkg/errors"
)

// QuestionDelete 代表获取问题详情
func (api *QAApi) QuestionDelete (c *gin.Context)  {
	qaService := c.MustMake(provider.QaKey).(provider.Service)
	type Param struct {
		ID int64 `json:"id" binding:"required"`
	}
	param := &Param{}
	if err := c.ShouldBind(param); err != nil {
		c.AbortWithError(404, err); return
	}

	question, err := qaService.GetQuestion(c, param.ID)
	if err != nil {
		c.AbortWithError(500, err); return
	}
	if question == nil {
		c.AbortWithError(500, errors.New("问题不存在")); return
	}

	user := auth.GetAuthUser(c)
	if user.ID != question.AuthorID {
		c.AbortWithError(500, errors.New("无权限操作")); return
	}

	ctx := provider.ContextWithUserID(c, user.ID)
	if err := qaService.DeleteQuestion(ctx, question.ID); err != nil {
		c.AbortWithError(500, err); return
	}
	c.ISetOkStatus().IJson(question)
}
