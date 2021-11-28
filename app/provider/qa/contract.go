package qa

import (
	"bbs/app/provider/user"
	"context"
	"time"
)

const QaKey = "qa"

// Service 代表qa的服务
type Service interface {

	// GetQuestions 获取问题列表，question简化结构
	GetQuestions(ctx context.Context, pager *Pager) ([]*Question, error)
	// GetQuestion 获取某个问题详情，question简化结构
	GetQuestion(ctx context.Context, questionID int64) (*Question, error)
	// PostQuestion 上传某个问题
	// ctx必须带操作人id
	PostQuestion(ctx context.Context, question *Question) error

	// QuestionLoadAuthor 问题加载Author字段
	QuestionLoadAuthor(ctx context.Context, question *Question) error
	// QuestionsLoadAuthor 批量加载Author字段
	QuestionsLoadAuthor(ctx context.Context, questions *[]*Question) error

	// QuestionLoadAnswers 单个问题加载Answers
	QuestionLoadAnswers(ctx context.Context, question *Question) error
	// QuestionsLoadAnswers 批量问题加载Answers
	QuestionsLoadAnswers(ctx context.Context, questions *[]*Question) error

	// PostAnswer 上传某个回答
	// ctx必须带操作人信息
	PostAnswer(ctx context.Context, answer *Answer) error
	// GetAnswer 获取回答
	GetAnswer(ctx context.Context, answerID int64) (*Answer, error)

	// DeleteQuestion 删除问题，同时删除对应的回答
	// ctx必须带操作人信息
	DeleteQuestion(ctx context.Context, questionID int64) error
	// DeleteAnswer 删除某个回答
	// ctx必须带操作人信息
	DeleteAnswer(ctx context.Context, answerID int64) error

	// UpdateQuestion 代表更新问题, 只会对比其中的context，title两个字段，其他字段不会对比
	// ctx必须带操作人
	UpdateQuestion(ctx context.Context, question *Question) error
}

// Question 代表问题
type Question struct {
	ID        int64     `gorm:"column:id;primaryKey"`
	Title     string    `gorm:"column:title"`
	Context   string    `gorm:"column:context"`
	AuthorID  int64     `gorm:"column:author_id"`
	AnswerNum int       `gorm:"column:answer_num"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	Author  *user.User `gorm:"foreignKey:AuthorID"`
	Answers []*Answer  `gorm:"foreignKey:QuestionID"`
}

// Answer 代表一个回答
type Answer struct {
	ID         int64     `gorm:"column:id;primaryKey"`
	QuestionID int64     `gorm:"column:question_id"`
	Content    string    `gorm:"column:context"`
	AuthorID   int64     `gorm:"column:author_id"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`

	Author   *user.User `gorm:"foreignKey:author_id"`
	Question *Question  `gorm:"foreignKey:question_id"`
}

// Pager 代表分页极致
type Pager struct {
	Total int
	Start int
	Size  int
}
