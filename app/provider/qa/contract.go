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
    GetQuestion(ctx context.Context, questionID int64)(*Question, error)
    // PostQuestion 上传某个问题
    // ctx必须带操作人id
    PostQuestion(ctx context.Context, question *Question) error

    // QuestionLoadAuthor 问题加载Author字段
    QuestionLoadAuthor(ctx context.Context, question *Question) error
    // QuestionsLoadAuthor 批量加载Author字段
    QuestionsLoadAuthor(ctx context.Context, questions []*Question) error

    // QuestionLoadAnswers 单个问题加载Answers
    QuestionLoadAnswers(ctx context.Context, question *Question) error
    // QuestionsLoadAnswers 批量问题加载Answers
    QuestionsLoadAnswers(ctx context.Context, questions []*Question) error

    // PostAnswer 上传某个回答
    // ctx必须带操作人信息
    PostAnswer(ctx context.Context, answer *Answer) error

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

// ContextWithUserID 将userID传入到context中，一些必要的操作必须进行userID验证
func ContextWithUserID(ctx context.Context, userID int64) context.Context {
    return context.WithValue(ctx, "qa-op-user-id", userID)
}

// ParseUserIDFromContext 从context中获取userID, 如果没有就返回0
func ParseUserIDFromContext(ctx context.Context) int64 {
    userID := ctx.Value("qa-op-user-id")
    if userID == nil {
        return 0
    }
    return userID.(int64)
}

// Question 代表问题
type Question struct {
    ID int64
    Title string
    Context string
    AuthorID int64
    CreatedAt time.Time
    UpdatedAt time.Time

    Author *user.User
    Answers []*Answer
}

type Answer struct {
    ID int64
    QuestionID int64
    Content string
    ParentID int64
    AuthorID int64
    CreatedAt time.Time
    UpdatedAt time.Time

    Author *user.User
}

type Pager struct {
    Total int
    Start int
    Size int
}
