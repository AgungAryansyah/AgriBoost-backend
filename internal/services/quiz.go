package services

import (
	"AgriBoost/internal/models/dto"
	entity "AgriBoost/internal/models/entities"
	"AgriBoost/internal/repositories"
	"context"
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QuizServiceItf interface {
	GetAllQuizzes(quiz *[]entity.Quiz) error
	GetQuizWithQuestionAndOption(quizDto *dto.QuizDto, quizParam dto.QuizParam) error
	CreateAttempt(answers dto.UserAnswersDto) error
}

type QuizService struct {
	quizRepo repositories.QuizRepoItf
	userRepo repositories.UserRepoItf
}

func NewQuizService(quizRepo repositories.QuizRepoItf, userRepo repositories.UserRepoItf) QuizServiceItf {
	return &QuizService{
		quizRepo: quizRepo,
		userRepo: userRepo,
	}
}

func (q *QuizService) GetAllQuizzes(quiz *[]entity.Quiz) error {
	return q.quizRepo.GetAllQuizzes(quiz)
}

func (q *QuizService) GetQuizWithQuestionAndOption(quizDto *dto.QuizDto, quizParam dto.QuizParam) error {
	return q.quizRepo.GetQuizWithQuestionAndOption(quizDto, quizParam)
}

func (q *QuizService) CreateAttempt(answers dto.UserAnswersDto) error {
	var (
		score       int
		mutex       sync.Mutex
		wg          sync.WaitGroup
		errChan     = make(chan error, 1)
		ctx, cancel = context.WithCancel(context.Background())
	)
	defer cancel()

	for questionId, answerId := range answers.Answers {
		wg.Add(1)
		go func(qid, aid uuid.UUID) {
			defer wg.Done()

			select {
			case <-ctx.Done():
				return
			default:
			}

			var answer entity.QuestionOption
			if err := q.quizRepo.GetCorrectOption(&answer, qid); err != nil {
				errChan <- err
				cancel()
				return
			}

			var question entity.Question
			if err := q.quizRepo.GetQuestion(&question, qid); err != nil {
				errChan <- err
				cancel()
				return
			}

			if question.QuizId != answers.QuizId {
				errChan <- errors.New("invalid answer for question")
				cancel()
				return
			}

			if answer.OptionId == aid {
				mutex.Lock()
				score += question.Score
				mutex.Unlock()
			}
		}(questionId, answerId)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	attempt := entity.QuizAttempt{
		AttemptId:    uuid.New(),
		UserId:       answers.UserId,
		QuizId:       answers.QuizId,
		TotalScore:   score,
		FinishedTime: time.Now(),
	}

	var bestAttempt entity.QuizAttempt
	if err := q.quizRepo.GetBestAttempt(&bestAttempt, attempt.UserId, attempt.QuizId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			q.userRepo.AddQuizPoint(dto.UserParam{Id: attempt.UserId}, attempt.TotalScore)
		} else {
			return err
		}
	} else if bestAttempt.TotalScore < score {
		q.userRepo.AddQuizPoint(dto.UserParam{Id: attempt.UserId}, score-bestAttempt.TotalScore)
	}

	return q.quizRepo.CreteAttempt(&attempt)
}
