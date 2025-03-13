package dto

import (
	entity "AgriBoost/internal/models/entities"
	"sync"

	"github.com/google/uuid"
)

type QuizParam struct {
	QuizId uuid.UUID `json:"quiz_id" validate:"required,uuid"`
}

type CreateAttempt struct {
	UserId     uuid.UUID `json:"user_id" validate:"required,uuid"`
	QuizId     uuid.UUID `json:"quiz_id" validate:"required,uuid"`
	TotalScore int       `json:"total_score" validate:"required,min=1"`
}

type QuizDto struct {
	QuizId    uuid.UUID     `json:"quiz_id"`
	Theme     string        `json:"theme"`
	Title     string        `json:"title"`
	Questions []QuestionDto `json:"questions"`
}

type QuestionDto struct {
	QuestionId    uuid.UUID   `json:"question_id"`
	QuestionText  string      `json:"question_text"`
	QuestionImage string      `json:"question_image"`
	Options       []OptionDto `json:"options"`
}

type OptionDto struct {
	OptionId    uuid.UUID `json:"option_id"`
	OptionText  string    `json:"option_text"`
	OptionImage string    `json:"option_image"`
}

type UserAnswersDto struct {
	UserId  uuid.UUID               `json:"user_id" validate:"required,uuid"`
	QuizId  uuid.UUID               `json:"quiz_id" validate:"required,uuid"`
	Answers map[uuid.UUID]uuid.UUID `json:"answers" validate:"required,answers_map"`
}

func QuizWithOptionAndoptionToDto(quiz entity.Quiz, quizDto *QuizDto) {
	*quizDto = QuizDto{
		QuizId:    quiz.QuizId,
		Theme:     quiz.Theme,
		Title:     quiz.Title,
		Questions: make([]QuestionDto, len(quiz.Questions)),
	}

	var wg sync.WaitGroup
	mu := sync.Mutex{}

	for i, question := range quiz.Questions {
		wg.Add(1)

		go func(i int, question entity.Question) {
			defer wg.Done()

			questionDto := QuestionDto{
				QuestionId:    question.QuestionId,
				QuestionText:  question.QuestionText,
				QuestionImage: question.QuestionImage,
				Options:       make([]OptionDto, len(question.Options)),
			}

			for j, option := range question.Options {
				questionDto.Options[j] = OptionDto{
					OptionId:    option.OptionId,
					OptionText:  option.OptionText,
					OptionImage: option.OptionImage,
				}
			}

			mu.Lock()
			quizDto.Questions[i] = questionDto
			mu.Unlock()
		}(i, question)
	}

	wg.Wait()
}
