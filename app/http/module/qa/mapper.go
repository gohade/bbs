package qa

import (
	"bbs/app/http/module/user"
	"bbs/app/provider/qa"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func ConvertAnswerToDTO(answer *qa.Answer) *AnswerDTO {
	if answer == nil {
		return nil
	}
	author := user.ConvertUserToDTO(answer.Author)
	if author == nil {
		author = &user.UserDTO{
			ID: answer.AuthorID,
		}
	}
	return &AnswerDTO{
		ID:        answer.ID,
		Content:   answer.Context,
		CreatedAt: answer.CreatedAt,
		UpdatedAt: answer.UpdatedAt,
		Author:    author,
	}
}

// ConvertAnswersToDTO 将answers转化为带有tree结构的AnswerDTO
func ConvertAnswersToDTO(answers []*qa.Answer) []*AnswerDTO {
	if answers == nil {
		return nil
	}

	ret := make([]*AnswerDTO, 0, len(answers))
	for _, answer := range answers {
		ret = append(ret, ConvertAnswerToDTO(answer))
	}
	return ret
}

// ConvertQuestionToDTO 将question转换为DTO
func ConvertQuestionToDTO(question *qa.Question, flags map[string]string) *QuestionDTO {
	if question == nil {
		return nil
	}
	author := user.ConvertUserToDTO(question.Author)
	if author == nil {
		author = &user.UserDTO{
			ID: question.AuthorID,
		}
	}

	context := question.Context
	if flags != nil {
		if isShortContext, ok := flags["is_short_context"]; ok && isShortContext == "true" {
			context = getShortContext(context)
		}
	}

	return &QuestionDTO{
		ID:        question.ID,
		Title:     question.Title,
		Context:   context,
		CreatedAt: question.CreatedAt,
		UpdatedAt: question.UpdatedAt,
		Author:    author,
		Answers:   ConvertAnswersToDTO(question.Answers),
	}
}

func getShortContext(context string) string {
	p := strings.NewReader(context)
	doc, _ := goquery.NewDocumentFromReader(p)

	doc.Find("script").Each(func(i int, el *goquery.Selection) {
		el.Remove()
	})

	text := doc.Text()
	if len(text) > 20 {
		text = text[:20] + "..."
	}
	return text
}

// ConvertQuestionsToDTO 将questions转换为DTO
func ConvertQuestionsToDTO(questions []*qa.Question) []*QuestionDTO {
	if questions == nil {
		return nil
	}
	ret := make([]*QuestionDTO, 0, len(questions))
	for _, question := range questions {
		ret = append(ret, ConvertQuestionToDTO(question, map[string]string{"is_short_context": "true"}))
	}
	return ret
}
