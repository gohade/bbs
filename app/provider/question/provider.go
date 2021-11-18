package question

import (
	"github.com/gohade/hade/framework"
)

type QuestionProvider struct {
	framework.ServiceProvider

	c framework.Container
}

func (sp *QuestionProvider) Name() string {
	return QuestionKey
}

func (sp *QuestionProvider) Register(c framework.Container) framework.NewInstance {
	return NewQuestionService
}

func (sp *QuestionProvider) IsDefer() bool {
	return false
}

func (sp *QuestionProvider) Params(c framework.Container) []interface{} {
	return []interface{}{c}
}

func (sp *QuestionProvider) Boot(c framework.Container) error {
	return nil
}

