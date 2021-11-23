package qa

import (
	"bbs/app/http/module/user"
	"bbs/app/provider/qa"
)

func ConvertAnswerToDTO(answer *qa.Answer) *AnswerDTO {
	if answer == nil {
		return nil
	}
	author := user.ConvertUserToDTO(answer.Author)
	if author == nil {
		author = &user.UserDTO{
			ID:        answer.AuthorID,
		}
	}
	return &AnswerDTO{
		ID:        answer.ID,
		Content:   answer.Content,
		CreatedAt: answer.CreatedAt,
		UpdatedAt: answer.UpdatedAt,
		Author:   author,
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
func ConvertQuestionToDTO(question *qa.Question) *QuestionDTO {
	if question == nil {
		return nil
	}
	author := user.ConvertUserToDTO(question.Author)
	if author == nil {
		author = &user.UserDTO{
			ID:        question.AuthorID,
		}
	}
	return &QuestionDTO{
		ID:        question.ID,
		Title:     question.Title,
		Context:   question.Context,
		CreatedAt: question.CreatedAt,
		UpdatedAt: question.UpdatedAt,
		Author:    author,
		Answers:   ConvertAnswersToDTO(question.Answers),
	}
}

// ConvertQuestionsToDTO 将questions转换为DTO
func ConvertQuestionsToDTO(questions []*qa.Question) []*QuestionDTO {
	if questions == nil {
		return nil
	}
	ret := make([]*QuestionDTO, 0, len(questions))
	for _, question := range questions {
		ret = append(ret, ConvertQuestionToDTO(question))
	}
	return ret
}

