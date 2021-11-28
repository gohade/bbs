package qa

import (
	"context"
	"fmt"
	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/contract"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
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
	question := &Question{}
	if err := q.ormDB.WithContext(ctx).First(question, questionID).Error; err != nil {
		return nil, err
	}
	return question, nil
}

func (q *QaService) PostQuestion(ctx context.Context, question *Question) error {
	if err := q.ormDB.WithContext(ctx).Create(question).Error; err != nil {
		return err
	}
	return nil
}

func (q *QaService) QuestionLoadAuthor(ctx context.Context, question *Question) error {
	if err := q.ormDB.WithContext(ctx).Preload("Author").First(question).Error; err != nil {
		return err
	}
	return nil
}

func (q *QaService) QuestionsLoadAuthor(ctx context.Context, questions *[]*Question) error {
	if err := q.ormDB.WithContext(ctx).Preload("Author").Find(questions).Error; err != nil {
		return err
	}
	return nil
}

func (q *QaService) QuestionLoadAnswers(ctx context.Context, question *Question) error {
	if err := q.ormDB.WithContext(ctx).Preload("Answers").First(question).Error; err != nil {
		return err
	}
	return nil
}

func (q *QaService) QuestionsLoadAnswers(ctx context.Context, questions *[]*Question) error {
	if err := q.ormDB.WithContext(ctx).Preload("Answers").Find(questions).Error; err != nil {
		return err
	}
	return nil
}

func (q *QaService) PostAnswer(ctx context.Context, answer *Answer) error {
	if answer.QuestionID == 0 {
		return errors.New("问题不存在")
	}
	question := &Question{ID: answer.QuestionID}
	if err := q.ormDB.WithContext(ctx).First(question).Error; err != nil {
		return err
	}
	if err := q.ormDB.WithContext(ctx).Create(answer).Error; err != nil {
		return err
	}
	question.AnswerNum = question.AnswerNum + 1
	if err := q.ormDB.WithContext(ctx).Save(question).Error; err != nil {
		return err
	}
	return nil
}

func (q *QaService) GetAnswer(ctx context.Context, answerID int64) (*Answer, error) {
	answer := &Answer{ID: answerID}
	if err := q.ormDB.WithContext(ctx).First(answer).Error; err != nil {
		return nil, err
	}
	return answer, nil
}

func (q *QaService) DeleteQuestion(ctx context.Context, questionID int64) error {
	question := &Question{ID: questionID}
	if err := q.ormDB.WithContext(ctx).Delete(question).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (q *QaService) DeleteAnswer(ctx context.Context, answerID int64) error {
	answer := &Answer{ID: answerID}
	if err := q.ormDB.WithContext(ctx).Delete(answer).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (q *QaService) UpdateQuestion(ctx context.Context, question *Question) error {
	questionDB := &Question{ID: question.ID}
	if err := q.ormDB.WithContext(ctx).First(questionDB).Error; err != nil {
		return errors.WithStack(err)
	}

	questionDB.UpdatedAt = time.Now()
	if question.Title != "" {
		questionDB.Title = question.Title
	}
	if question.Context != "" {
		questionDB.Context = question.Context
	}
	if err := q.ormDB.WithContext(ctx).Save(questionDB).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
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
