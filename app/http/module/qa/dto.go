package qa

import (
	"bbs/app/http/module/user"
	"time"
)

// QuestionDTO 问题列表返回结构
type QuestionDTO struct {
	ID        int64     `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Context   string    `json:"context,omitempty"` // 在列表页，只显示前200个字符
	AnswerNum int       `json:"answer_num"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Author  *user.UserDTO `json:"author,omitempty"`  // 作者
	Answers []*AnswerDTO  `json:"answers,omitempty"` // 回答
}

// AnswerDTO 回答返回结构
type AnswerDTO struct {
	ID        int64     `json:"id,omitempty"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Author *user.UserDTO `json:"author,omitempty"` // 作者
}
