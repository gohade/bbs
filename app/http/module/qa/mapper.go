package qa

import (
	"bbs/app/http/module/user"
	"bbs/app/provider/qa"
	"github.com/jianfengye/collection"
)

// 获取answer的树形结构
func getAnswerChildren(dto *AnswerDTO, answers []*qa.Answer) {
	if dto == nil {
		return
	}
	for _, answer := range answers {
		if dto.ID == answer.ParentID {
			if dto.Children == nil {
				dto.Children = []*AnswerDTO{}
			}

			childAnswerDTO :=  &AnswerDTO{
				ID:        answer.ID,
				Content:   answer.Content,
				AuthorID:  answer.AuthorID,
				CreatedAt: answer.CreatedAt,
				UpdatedAt: answer.UpdatedAt,
				Author:    user.ConvertUserToDTO(answer.Author),
				Children:  nil,
			}

			getAnswerChildren(childAnswerDTO, answers)

			dto.Children = append(dto.Children, childAnswerDTO)
		}
	}


	if len(dto.Children) > 0 {
		childColl := collection.NewObjPointCollection(dto.Children)
		objs := []*AnswerDTO{}
		childColl.SortByDesc("UpdatedAt").ToObjs(&objs)
		dto.Children = objs
	}
	return
}

// ConvertAnswersToDTO 将answers转化为带有tree结构的AnswerDTO
func ConvertAnswersToDTO(answers []*qa.Answer) []*AnswerDTO {
	if answers == nil {
		return nil
	}

	answerZero := &AnswerDTO{
		ID:        0,
		Children:  nil,
	}
	getAnswerChildren(answerZero, answers)
	return answerZero.Children
}

// ConvertQuestionToDTO 将question转换为DTO
func ConvertQuestionToDTO(question *qa.Question) *QuestionDTO {
	if question == nil {
		return nil
	}
	return &QuestionDTO{
		ID:        question.ID,
		Title:     question.Title,
		Context:   question.Context,
		AuthorID:  question.AuthorID,
		CreatedAt: question.CreatedAt,
		UpdatedAt: question.UpdatedAt,
		Author:    user.ConvertUserToDTO(question.Author),
		Answers:   ConvertAnswersToDTO(question.Answers),
	}
}
