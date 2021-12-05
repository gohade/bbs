package qa

import (
	"bbs/app/http/middleware/auth"
	provider "bbs/app/provider/qa"
	"github.com/gohade/hade/framework/gin"
)

type questionCreateParam struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// QuestionCreate 代表创建问题
// @Summary 创建问题
// @Description 创建问题
// @Accept  json
// @Produce  json
// @Tags qa
// @Param questionCreateParam body questionCreateParam true "创建问题参数"
// @Success 200 string Msg "操作成功"
// @Router /question/create [post]
func (api *QAApi) QuestionCreate(c *gin.Context) {
	qaService := c.MustMake(provider.QaKey).(provider.Service)

	param := &questionCreateParam{}
	if err := c.ShouldBind(param); err != nil {
		c.ISetStatus(400).IText(err.Error())
		return
	}

	user := auth.GetAuthUser(c)
	if user == nil {
		c.ISetStatus(500).IText("无权限操作")
		return
	}

	question := &provider.Question{
		Title:    param.Title,
		Context:  param.Content,
		AuthorID: user.ID,
	}
	if err := qaService.PostQuestion(c, question); err != nil {
		c.ISetStatus(500).IText(err.Error())
		return
	}

	c.ISetOkStatus().IText("操作成功")
}
