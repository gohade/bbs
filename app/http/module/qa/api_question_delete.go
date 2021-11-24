package qa

import (
	"bbs/app/http/middleware/auth"
	provider "bbs/app/provider/qa"
	"github.com/gohade/hade/framework/gin"
)

// QuestionDelete 删除问题
// @Summary 删除问题
// @Description 删除问题，同时删除问题中的所有答案
// @Accept  json
// @Produce  json
// @Tags qa
// @Param id query int true "删除id"
// @Success 200 {string} Msg "操作成功"
// @Router /question/delete [get]
func (api *QAApi) QuestionDelete(c *gin.Context) {
	qaService := c.MustMake(provider.QaKey).(provider.Service)
	id, exist := c.DefaultQueryInt64("id", 0)
	if !exist {
		c.ISetStatus(404).IText("参数错误")
		return
	}

	question, err := qaService.GetQuestion(c, id)
	if err != nil {
		c.ISetStatus(500).IText(err.Error())
		return
	}
	if question == nil {
		c.ISetStatus(500).IText("问题不存在")
		return
	}

	user := auth.GetAuthUser(c)
	if user.ID != question.AuthorID {
		c.ISetStatus(500).IText("无权限操作")
		return
	}

	ctx := provider.ContextWithUserID(c, user.ID)
	if err := qaService.DeleteQuestion(ctx, question.ID); err != nil {
		c.ISetStatus(500).IText(err.Error())
		return
	}
	c.ISetOkStatus().IText("操作成功")
}
