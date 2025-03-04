package repositories

import (
	"AgriBoost/internal/models/dto"
	entity "AgriBoost/internal/models/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QuizRepoItf interface {
	GetAllQuizzes(quiz *[]entity.Quiz) error
	GetQuizWithQuestionAndOption(quizDto *dto.QuizDto, quizParam dto.QuizParam) error
	GetCorrectOption(correctOption *entity.QuestionOption, questionId uuid.UUID) error
	CreteAttempt(attempt *entity.QuizAttempt) error
	GetQuestion(question *entity.Question, questionId uuid.UUID) error
}

type QuizRepo struct {
	db *gorm.DB
}

func NewQuizRepo(db *gorm.DB) QuizRepoItf {
	return &QuizRepo{
		db: db,
	}
}

func (q *QuizRepo) GetAllQuizzes(quiz *[]entity.Quiz) error {
	return q.db.Find(quiz).Error
}

func (q *QuizRepo) GetQuizWithQuestionAndOption(quizDto *dto.QuizDto, quizParam dto.QuizParam) error {
	var quiz entity.Quiz
	err := q.db.Preload("Question.option").First(&quiz, quizParam).Error

	if err != nil {
		return nil
	}
	dto.QuizWithOptionAndoptionToDto(quiz, quizDto)

	return err
}

func (q *QuizRepo) GetQuiz(quiz *entity.Quiz, quizParam dto.QuizParam) error {
	return q.db.First(&quiz, quizParam).Error
}

func (q *QuizRepo) CreteAttempt(attempt *entity.QuizAttempt) error {
	return q.db.Create(attempt).Error
}

func (q *QuizRepo) GetCorrectOption(correctOption *entity.QuestionOption, questionId uuid.UUID) error {
	return q.db.Where("question_id = ? AND is correct = ?", questionId, true).First(&correctOption).Error
}

func (q *QuizRepo) GetQuestion(question *entity.Question, questionId uuid.UUID) error {
	return q.db.First(question, questionId).Error
}
