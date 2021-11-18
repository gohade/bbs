package question

import "github.com/gohade/hade/framework"

type QuestionService struct {
	container framework.Container
}

func NewQuestionService(params ...interface{}) (interface{}, error) {
	container := params[0].(framework.Container)
	return &QuestionService{container: container}, nil
}

func (s *QuestionService) Foo() string {
    return ""
}
