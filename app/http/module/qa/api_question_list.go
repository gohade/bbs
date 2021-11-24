package qa

import (
	provider "bbs/app/provider/qa"
	"github.com/gohade/hade/framework/gin"
)

// QuestionList 获取问题列表
// @Summary 获取问题列表
// @Description 获取问题列表，包含作者信息，不包含回答
// @Accept  json
// @Produce  json
// @Tags qa
// @Param page query int false "列表页页数"
// @Param size query int false "列表页单页个数"
// @Success 200 []QuestionDTO questions "问题列表"
// @Router /question/list [get]
func (api *QAApi) QuestionList(c *gin.Context) {
	qaService := c.MustMake(provider.QaKey).(provider.Service)
	page, _ := c.DefaultQueryInt("page", 1)
	size, _ := c.DefaultQueryInt("size", 10)

	start := (page - 1) * size
	pager := provider.Pager{
		Start: start,
		Size:  size,
	}
	questions, err := qaService.GetQuestions(c, &pager)
	if err != nil {
		c.ISetStatus(500).IText(err.Error())
		return
	}

	if err := qaService.QuestionsLoadAuthor(c, questions); err != nil {
		c.ISetStatus(500).IText(err.Error())
		return
	}

	questionsDTO := ConvertQuestionsToDTO(questions)

	c.ISetOkStatus().IJson(questionsDTO)
}
