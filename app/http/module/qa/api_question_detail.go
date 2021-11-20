package qa

import (
	provider "bbs/app/provider/qa"
	"github.com/gohade/hade/framework/gin"
)

// QuestionDetail 代表获取问题详情
func (api *QAApi) QuestionDetail(c *gin.Context)  {
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

	if err := qaService.QuestionLoadAuthor(c, question); err != nil {
		c.AbortWithError(500, err); return
	}
	if err := qaService.QuestionLoadAnswers(c, question); err != nil {
		c.AbortWithError(500, err); return
	}

	c.ISetOkStatus().IJson(question)
}
