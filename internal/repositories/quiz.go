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
	GetBestAttempt(attempt *entity.QuizAttempt, userId uuid.UUID, quizId uuid.UUID) error
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

	if err := q.db.Preload("Questions.Options").Find(&quiz, quizParam).Error; err != nil {
		return err
	}

	dto.QuizWithOptionAndoptionToDto(quiz, quizDto)

	return nil
}

func (q *QuizRepo) GetQuiz(quiz *entity.Quiz, quizParam dto.QuizParam) error {
	return q.db.First(&quiz, quizParam).Error
}

func (q *QuizRepo) CreteAttempt(attempt *entity.QuizAttempt) error {
	return q.db.Create(attempt).Error
}

func (q *QuizRepo) GetCorrectOption(correctOption *entity.QuestionOption, questionId uuid.UUID) error {
	return q.db.Where("question_id = ? AND is_correct = ?", questionId, true).First(&correctOption).Error
}

func (q *QuizRepo) GetQuestion(question *entity.Question, questionId uuid.UUID) error {
	return q.db.First(question, questionId).Error
}

func (q *QuizRepo) GetBestAttempt(attempt *entity.QuizAttempt, userId uuid.UUID, quizId uuid.UUID) error {
	var count int64
	q.db.Model(&attempt).Where("user_id = ? AND quiz_id = ?", userId, quizId).Count(&count)

	if count == 0 {
		return gorm.ErrRecordNotFound
	}
	return q.db.Where("user_id = ? AND quiz_id = ?", userId, quizId).Order("total_score DESC, attempt_id").First(attempt).Error
}
