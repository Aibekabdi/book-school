package service

import (
	"book-school/internal/models"
	"book-school/internal/repository"
	"book-school/pkg/tts"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/sync/errgroup"
)

type TestService struct {
	testRepo     repository.Test
	questionRepo repository.Question
	answerRepo   repository.Answer
	studentRepo  repository.Student
	classRepo    repository.Class
	bookRepo     repository.Book

	imageUrl string
}

func newTestService(testRepo repository.Test, questionRepo repository.Question, answerRepo repository.Answer, studentRepo repository.Student, classRepo repository.Class, bookRepo repository.Book, imageUrl string) *TestService {
	return &TestService{
		testRepo:     testRepo,
		questionRepo: questionRepo,
		answerRepo:   answerRepo,
		studentRepo:  studentRepo,
		classRepo:    classRepo,
		bookRepo:     bookRepo,
		imageUrl:     imageUrl,
	}
}

func (s *TestService) validate(test models.Test) error {
	hasOneCorrect := func(arr []models.Answer) bool {
		for _, a := range arr {
			if a.Correct {
				return true
			}
		}
		return false
	}

	if test.BookId == 0 {
		return fmt.Errorf("validate: empty book id")
	}

	if len(test.Questions) == 0 {
		return fmt.Errorf("validate: empty questions")
	}

	for _, q := range test.Questions {
		if strings.ReplaceAll(q.Question, " ", "") == "" {
			return fmt.Errorf("validate: empty question")
		}
		if len(q.Answers) == 0 {
			return fmt.Errorf("validate: empty answers")
		}
		if len(q.Answers) > 5 {
			return fmt.Errorf("validate: too many answers")
		}
		if len(q.Answers) < 3 {
			return fmt.Errorf("validate: not enough answers")
		}
		for _, a := range q.Answers {
			if strings.ReplaceAll(a.Answer, " ", "") == "" {
				return fmt.Errorf("validate: empty answer")
			}
		}
		if !hasOneCorrect(q.Answers) {
			return fmt.Errorf("validate: all answers false")
		}
	}

	return nil
}

func (s *TestService) RePass(ctx context.Context, user models.User, testId int) error {
	test, err := s.testRepo.GetCompleteTest(context.WithValue(ctx, models.CompleteTestId, uint(testId)))
	if err != nil {
		return err
	}

	err = s.testRepo.DeleteSavedTest(ctx, uint(test.Id))
	if err != nil {
		return err
	}

	err = s.studentRepo.TakePoints(ctx, user.Id, test.Points)
	if err != nil {
		return err
	}

	return nil
}

func (s *TestService) Delete(ctx context.Context, testId uint) error {
	err := os.RemoveAll(fmt.Sprintf("./static/tests/%d/", testId))
	if err != nil {
		return fmt.Errorf("test service: delete: %w", err)
	}

	err = s.testRepo.Delete(ctx, testId)
	if err != nil {
		return fmt.Errorf("test service: delete: %w", err)
	}

	return nil
}

func (s *TestService) Create(ctx context.Context, test models.Test, questions, answers []*multipart.FileHeader) error {
	var err error

	if err = s.validate(test); err != nil {
		return fmt.Errorf("test service: create: %w", err)
	}

	t, err := s.testRepo.Get(context.WithValue(ctx, models.BookId, test.BookId))
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("test service: create: %w", err)
		}
	}
	if len(t) > 0 {
		return fmt.Errorf("test service: create: test for book exist")
	}

	book, err := s.bookRepo.Get(context.WithValue(ctx, models.BookId, test.BookId))
	if err != nil {
		return fmt.Errorf("test service: create: %w", err)
	}

	test.Lang = book[0].Language

	test.Id, err = s.testRepo.Create(ctx, test.BookId, test.Lang)
	if err != nil {
		return fmt.Errorf("test service: create: %w", err)
	}

	isKaz := true
	if test.Lang == "RU" {
		isKaz = false
	}

	errs, _ := errgroup.WithContext(context.Background())
	var mu sync.Mutex

	for i, q := range test.Questions {
		func(i int, q models.Question) {
			errs.Go(func() error {
				mu.Lock()
				defer mu.Unlock()
				audio, err := tts.TextToSpeech(q.Question, fmt.Sprintf("./static/tests/%d/%d/", test.Id, i+1), i+1, isKaz)
				if err != nil {
					return fmt.Errorf("test service: create: %w", err)
				}

				test.Questions[i].Audio = s.imageUrl + audio[1:]
				return nil
			})
		}(i, q)

		for j, a := range q.Answers {
			func(i, j int, a models.Answer) {
				errs.Go(func() error {
					mu.Lock()
					defer mu.Unlock()
					audio, err := tts.TextToSpeech(a.Answer, fmt.Sprintf("./static/tests/%d/%d/answers/", test.Id, i+1), j+1, isKaz)
					if err != nil {
						return fmt.Errorf("test service: create: %w", err)
					}

					test.Questions[i].Answers[j].Audio = s.imageUrl + audio[1:]
					return nil
				})
			}(i, j, a)
		}
	}

	err = errs.Wait()
	if err != nil {
		if err = s.Delete(ctx, test.Id); err != nil {
			return fmt.Errorf("test service: create: %w", err)
		}
		return err
	}

	for i, question := range test.Questions {
		question.TestId = test.Id

		if question.WithImage {
			image := questions[0]
			path, err := saveFiles(fmt.Sprintf("./static/tests/%d/", test.Id), strconv.Itoa(i+1), image)
			if err != nil {
				if err = s.Delete(ctx, test.Id); err != nil {
					return fmt.Errorf("test service: create: %w", err)
				}
				return fmt.Errorf("test service: create: %w", err)
			}
			question.Image = s.imageUrl + strings.ReplaceAll(path, "./static/", "/static/")
			questions = questions[1:]
		}

		question.Id, err = s.questionRepo.Create(ctx, question)
		if err != nil {
			if err = s.Delete(ctx, test.Id); err != nil {
				return fmt.Errorf("test service: create: %w", err)
			}
			return fmt.Errorf("test service: create: %w", err)
		}

		for _, answer := range question.Answers {
			answer.QuestionId = question.Id

			if answer.WithImage {
				image := answers[0]
				path, err := saveFiles(fmt.Sprintf("./static/tests/%d/%d/", test.Id, i+1), "answers", image)
				if err != nil {
					if err = s.Delete(ctx, test.Id); err != nil {
						return fmt.Errorf("test service: create: %w", err)
					}
					return fmt.Errorf("test service: create: %w", err)
				}
				answer.Image = s.imageUrl + strings.ReplaceAll(path, "./static/", "/static/")
				answers = answers[1:]
			}

			if err = s.answerRepo.Create(ctx, answer); err != nil {
				if err = s.Delete(ctx, test.Id); err != nil {
					return fmt.Errorf("test service: create: %w", err)
				}
				return fmt.Errorf("test service: create: %w", err)
			}
		}
	}

	return nil
}

func (s *TestService) Get(ctx context.Context) ([]models.Test, error) {
	tests, err := s.testRepo.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("test serice: get test by book id: %w", err)
	}

	if len(tests) == 0 {
		return nil, fmt.Errorf("test serice: get test by book id: %w", sql.ErrNoRows)
	}

	user := ctx.Value(models.UserCtx).(models.User)
	if user.Role == models.StudentRole {
		newCtx := context.WithValue(ctx, models.StudentId, user.Id)
		newCtx = context.WithValue(newCtx, models.TestId, tests[0].Id)

		_, err := s.testRepo.GetCompleteTest(newCtx)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				for i, test := range tests {
					tests[i].Questions, err = s.questionRepo.GetByTestId(ctx, test.Id)
					if err != nil {
						return nil, fmt.Errorf("test serice: get test by book id: %w", err)
					}

					for j, question := range tests[i].Questions {
						tests[i].Questions[j].Answers, err = s.answerRepo.GetByQuestionId(ctx, question.Id)
						if err != nil {
							return nil, fmt.Errorf("test serice: get test by book id: %w", err)
						}
					}
				}

				return tests, nil
			}
			return nil, fmt.Errorf("test serice: get test by book id: %w", err)
		}
		return nil, nil
	}

	for i, test := range tests {
		tests[i].Questions, err = s.questionRepo.GetByTestId(ctx, test.Id)
		if err != nil {
			return nil, fmt.Errorf("test serice: get test by book id: %w", err)
		}

		for j, question := range tests[i].Questions {
			tests[i].Questions[j].Answers, err = s.answerRepo.GetByQuestionId(ctx, question.Id)
			if err != nil {
				return nil, fmt.Errorf("test serice: get test by book id: %w", err)
			}
		}
	}

	return tests, nil
}

func (s *TestService) GetCompleteTest(ctx context.Context, testId, studentId uint) (models.CompleteTest, error) {
	newCtx := context.WithValue(ctx, models.StudentId, studentId)
	newCtx = context.WithValue(newCtx, models.TestId, testId)

	test, err := s.testRepo.GetCompleteTest(newCtx)
	if err != nil {
		return models.CompleteTest{}, fmt.Errorf("test service: get complete test: %w", err)
	}

	test.Answers, err = s.answerRepo.GetCompleteAnswers(ctx, test.Id)
	if err != nil {
		return models.CompleteTest{}, fmt.Errorf("test service: get complete test: %w", err)
	}

	return test, nil
}

func (s *TestService) GetTestForTeacher(ctx context.Context, bookId, teacherId uint) ([]models.TestForTeacher, error) {
	tests, err := s.Get(context.WithValue(ctx, models.BookId, bookId))
	if err != nil {
		return nil, fmt.Errorf("test service: get test for teacher: %w", err)
	}

	if len(tests) == 0 {
		return nil, fmt.Errorf("test service: get test for teacher: invalid test")
	}

	classes, err := s.classRepo.GetAll(context.WithValue(ctx, models.TeacherId, teacherId))
	if err != nil {
		return nil, fmt.Errorf("test service: get test for teacher: %w", err)
	}

	all := []models.TestForTeacher{}

	for _, class := range classes {
		students, err := s.studentRepo.GetAll(context.WithValue(ctx, models.ClassId, class.Id))
		if err != nil {
			return nil, fmt.Errorf("test service: get test for teacher: %w", err)
		}

		for _, student := range students {
			curTest := models.Test{}
			curTest.Id = tests[0].Id
			curTest.BookId = tests[0].BookId
			curTest.Questions = make([]models.Question, len(tests[0].Questions))

			for i, q := range tests[0].Questions {
				curTest.Questions[i] = models.Question{
					Id:        q.Id,
					TestId:    q.TestId,
					WithImage: q.WithImage,
					Image:     q.Image,
					Question:  q.Question,
				}
				curTest.Questions[i].Answers = make([]models.Answer, len(q.Answers))
				for j, a := range q.Answers {
					curTest.Questions[i].Answers[j] = models.Answer{
						Id:               a.Id,
						QuestionId:       a.QuestionId,
						WithImage:        a.WithImage,
						Image:            a.Image,
						Answer:           a.Answer,
						Correct:          a.Correct,
						IsStudentsAnswer: a.IsStudentsAnswer,
					}
				}
			}

			test, err := s.GetCompleteTest(ctx, curTest.Id, student.Id)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					continue
				}
				return nil, fmt.Errorf("test service: get test for teacher: %w", err)
			}

			for i, q := range curTest.Questions {
			answers:
				for j, a := range q.Answers {
					for _, s := range test.Answers {
						if s.QuestionId == q.Id && s.AnswerId == a.Id {
							curTest.Questions[i].Answers[j].IsStudentsAnswer = true
							break answers
						}
					}
				}
			}

			testForTeacher := models.TestForTeacher{
				Student: student,
				Test:    curTest,
			}

			all = append(all, testForTeacher)
		}
	}

	return all, nil
}

func (s *TestService) CompleteTest(ctx context.Context, complteTest models.CompleteTest) (models.CompleteTestResp, error) {
	_, err := s.GetCompleteTest(ctx, complteTest.TestId, complteTest.StudentId)
	if err == nil {
		return models.CompleteTestResp{}, fmt.Errorf("test service: complete test: you passed test once")
	} else if !errors.Is(err, sql.ErrNoRows) {
		return models.CompleteTestResp{}, fmt.Errorf("test service: complete test: %w", err)
	}

	test, err := s.Get(context.WithValue(ctx, models.TestId, complteTest.TestId))
	if err != nil {
		return models.CompleteTestResp{}, fmt.Errorf("test service: comlete test: %w", err)
	}

	var resp models.CompleteTestResp

	resp.CorrectAnswers = make([]bool, len(test[0].Questions))

	complteTest.Points = 0
answers:
	for i, q := range test[0].Questions {
		for _, a := range q.Answers {
			for _, s := range complteTest.Answers {
				if s.QuestionId == q.Id && s.AnswerId == a.Id && a.Correct {
					complteTest.Points += models.PointsForAnswer
					resp.CorrectAnswers[i] = true
					resp.CorrectAnswersCount++
					continue answers
				}
			}
		}
		resp.CorrectAnswers[i] = false
	}

	resp.Points = int(complteTest.Points)

	if err = s.studentRepo.GivePoints(ctx, complteTest.StudentId, complteTest.Points); err != nil {
		return models.CompleteTestResp{}, fmt.Errorf("test service: complete test: %w", err)
	}

	id, err := s.testRepo.SaveTest(ctx, complteTest)
	if err != nil {
		if err = s.studentRepo.TakePoints(ctx, complteTest.StudentId, complteTest.Points); err != nil {
			return models.CompleteTestResp{}, fmt.Errorf("test service: complete test: %w", err)
		}

		return models.CompleteTestResp{}, fmt.Errorf("test service: complete test: %w", err)
	}

	resp.TestId = int(id)

	for _, answer := range complteTest.Answers {
		answer.CompleteTestId = id
		if err = s.answerRepo.SaveAnswers(ctx, answer); err != nil {
			if err = s.studentRepo.TakePoints(ctx, complteTest.StudentId, complteTest.Points); err != nil {
				return models.CompleteTestResp{}, fmt.Errorf("test service: complete test: %w", err)
			}

			if err = s.testRepo.DeleteSavedTest(ctx, id); err != nil {
				return models.CompleteTestResp{}, fmt.Errorf("test service: complete test: %w", err)
			}

			return models.CompleteTestResp{}, fmt.Errorf("test service: complete test: %w", err)
		}
	}

	return resp, nil
}
