package qa

import "github.com/gohade/hade/framework"

type QaService struct {
	container framework.Container
}

func NewQaService(params ...interface{}) (interface{}, error) {
	container := params[0].(framework.Container)
	return &QaService{container: container}, nil
}

func (s *QaService) Foo() string {
    return ""
}
