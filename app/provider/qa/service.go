package qa

import (
	"context"
	"fmt"
	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/contract"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type QaService struct {
	container framework.Container // 容器
	ormDB     *gorm.DB            // db
	logger    contract.Log        // log
}

func (q *QaService) GetQuestions(ctx context.Context, pager *Pager) ([]*Question, error) {
	questions := make([]*Question, 0, pager.Size)
	if err := q.ormDB.Order("created_at desc").Offset(pager.Start).Limit(pager.Size).Find(&questions).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return questions, nil
}

func (q *QaService) GetQuestion(ctx context.Context, questionID int64) (*Question, error) {
	panic("implement me")
}

func (q *QaService) PostQuestion(ctx context.Context, question *Question) error {
	panic("implement me")
}

func (q *QaService) QuestionLoadAuthor(ctx context.Context, question *Question) error {
	panic("implement me")
}

func (q *QaService) QuestionsLoadAuthor(ctx context.Context, questions []*Question) error {
	panic("implement me")
}

func (q *QaService) QuestionLoadAnswers(ctx context.Context, question *Question) error {
	panic("implement me")
}

func (q *QaService) QuestionsLoadAnswers(ctx context.Context, questions []*Question) error {
	panic("implement me")
}

func (q *QaService) PostAnswer(ctx context.Context, answer *Answer) error {
	panic("implement me")
}

func (q *QaService) GetAnswer(ctx context.Context, answerID int64) (*Answer, error) {
	panic("implement me")
}

func (q *QaService) DeleteQuestion(ctx context.Context, questionID int64) error {
	panic("implement me")
}

func (q *QaService) DeleteAnswer(ctx context.Context, answerID int64) error {
	panic("implement me")
}

func (q *QaService) UpdateQuestion(ctx context.Context, question *Question) error {
	panic("implement me")
}

func NewQaService(params ...interface{}) (interface{}, error) {
	container := params[0].(framework.Container)
	ormService := container.MustMake(contract.ORMKey).(contract.ORMService)
	logger := container.MustMake(contract.LogKey).(contract.Log)

	db, err := ormService.GetDB()
	if err != nil {
		logger.Error(context.Background(), "获取gormDB错误", map[string]interface{}{
			"err": fmt.Sprintf("%+v", err),
		})
		return nil, err
	}
	return &QaService{container: container, ormDB: db, logger: logger}, nil
}
