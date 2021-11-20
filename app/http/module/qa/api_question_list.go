package qa

import (
	provider "bbs/app/provider/qa"
	"github.com/gohade/hade/framework/gin"
)

// QuestionList 代表获取问题列表
func (api *QAApi) QuestionList(c *gin.Context)  {
	qaService := c.MustMake(provider.QaKey).(provider.Service)
	type Param struct {
		Page int `json:"page" binding:"required"`
		Size int `json:"size" binding:"required"`
	}
	param := &Param{}
	if err := c.ShouldBind(param); err != nil {
		c.AbortWithError(404, err); return
	}

	start := (param.Page - 1) * param.Size
	pager := provider.Pager{
		Start: start,
		Size:  param.Size,
	}
	questions, err := qaService.GetQuestions(c, &pager)
	if err != nil {
		c.AbortWithError(500, err); return
	}

	if err := qaService.QuestionsLoadAuthor(c, questions); err != nil {
		c.AbortWithError(500, err); return
	}
	if err := qaService.QuestionsLoadAnswers(c, questions); err != nil {
		c.AbortWithError(500, err); return
	}

	c.ISetOkStatus().IJson(questions)
}
