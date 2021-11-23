package qa

import (
	"bbs/app/http/module/user"
	"time"
)

// QuestionDTO 问题列表返回结构
type QuestionDTO struct {
	ID int64
	Title string
	Context string
	CreatedAt time.Time
	UpdatedAt time.Time

	Author *user.UserDTO
	Answers [] *AnswerDTO
}

// AnswerDTO 回答返回结构
type AnswerDTO struct {
	ID int64
	Content string
	CreatedAt time.Time
	UpdatedAt time.Time

	Author *user.UserDTO
}
