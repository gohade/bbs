package qa

import (
	provider "bbs/app/provider/qa"
	"github.com/gohade/hade/framework/gin"
)

// QuestionDetail 获取问题详情
// @Summary 获取问题详细
// @Description 获取问题详情，包括问题的所有回答
// @Accept  json
// @Produce  json
// @Tags qa
// @Param id query int true "问题id"
// @Success 200 QuestionDTO question "问题详情，带回答和作者"
// @Router /question/detail [get]
func (api *QAApi) QuestionDetail(c *gin.Context) {
	qaService := c.MustMake(provider.QaKey).(provider.Service)
	id, exist := c.DefaultQueryInt64("id", 0)
	if !exist {
		c.ISetStatus(400).IText("参数错误")
		return
	}

	// 获取问题详情
	question, err := qaService.GetQuestion(c, id)
	if err != nil {
		c.ISetStatus(500).IText(err.Error())
		return
	}

	// 加载问题作者
	if err := qaService.QuestionLoadAuthor(c, question); err != nil {
		c.ISetStatus(500).IText(err.Error())
		return
	}
	// 加载所有答案
	if err := qaService.QuestionLoadAnswers(c, question); err != nil {
		c.ISetStatus(500).IText(err.Error())
		return
	}
	// 加载答案的所有作者
	if err := qaService.AnswersLoadAuthor(c, &question.Answers); err != nil {
		c.ISetStatus(500).IText(err.Error())
		return
	}

	// 输出转换
	questionDTO := ConvertQuestionToDTO(question, nil)

	c.ISetOkStatus().IJson(questionDTO)
}
