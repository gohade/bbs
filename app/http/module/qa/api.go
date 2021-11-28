package qa

import (
	"bbs/app/http/middleware/auth"
	"bbs/app/provider/qa"
	"github.com/gohade/hade/framework/gin"
)

type QAApi struct{}

// RegisterRoutes 注册路由
func RegisterRoutes(r *gin.Engine) error {
	api := &QAApi{}
	if !r.IsBind(qa.QaKey) {
		r.Bind(&qa.QaProvider{})
	}
	questionApi := r.Group("/question", auth.AuthMiddleware())
	{
		// 问题列表
		questionApi.GET("/list", api.QuestionList)
		// 问题详情
		questionApi.GET("/detail", api.QuestionDetail)
		// 创建问题
		questionApi.POST("/create", api.QuestionCreate)
		// 删除问题
		questionApi.POST("/delete", api.QuestionDelete)
		// 更新问题
		questionApi.POST("/edit", api.QuestionEdit)
	}
	answerApi := r.Group("/answer", auth.AuthMiddleware())
	{
		// 创建回答
		answerApi.POST("/create", api.AnswerCreate)
		// 删除回答
		answerApi.POST("/delete", api.AnswerDelete)
	}

	return nil
}
