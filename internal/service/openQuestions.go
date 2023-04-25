package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"book-school/internal/models"
	"book-school/internal/repository"
	"book-school/pkg/tts"
	"book-school/pkg/utils"
)

type CreativeTaskService struct {
	CreativeTaskRepo repository.CreativeTask
	studentRepo      repository.Student
	classRepo        repository.Class
	teacherRepo      repository.Teacher
	schoolRepo       repository.School
	bookRepo         repository.Book
	imageUrl         string
}

func newCreativeTaskService(CreativeTask repository.CreativeTask, studentRepo repository.Student, classRepo repository.Class, teacherRepo repository.Teacher,
	schoolRepo repository.School, bookRepo repository.Book, imageUrl string,
) *CreativeTaskService {
	return &CreativeTaskService{
		CreativeTaskRepo: CreativeTask,
		studentRepo:      studentRepo,
		classRepo:        classRepo,
		teacherRepo:      teacherRepo,
		schoolRepo:       schoolRepo,
		bookRepo:         bookRepo,
		imageUrl:         imageUrl,
	}
}

func (g *CreativeTaskService) GetCreativeTask(user models.User, category string, isCreative bool) ([]models.CreativeTask, error) {
	return g.CreativeTaskRepo.GetCreativeTask(user, category, isCreative)
}

func (g *CreativeTaskService) CreatePassCreativeTask(answer models.PassCreativeTask, isCreative bool) error {
	if len(answer.Answer) > 5000 {
		return errors.New("answer's letter is over limitted")
	}
	if isCreative {
		strStudentId := strconv.Itoa(int(answer.StudentId))
		strBookId := strconv.Itoa(int(answer.BookId))
		path := "./static/open/" + strStudentId + "/" + strBookId + "/"
		if err := utils.CreateDirectory("./static/open/"); err != nil {
			return err
		}
		if err := utils.CreateDirectory("./static/open/" + strStudentId); err != nil {
			return err
		}
		if err := utils.CreateDirectory(path); err != nil {
			return err
		}

		if err := savePreview(path+strconv.Itoa(int(answer.QuestionId)), answer.Img); err != nil {
			return fmt.Errorf("CreatePassCreativeTask service: cannot save image: %w", err)
		}
		answer.Answer = g.imageUrl + strings.ReplaceAll(path+strconv.Itoa(int(answer.QuestionId)), `./`, "/")
	}

	return g.CreativeTaskRepo.CreatePassCreativeTask(answer, isCreative)
}

func (g *CreativeTaskService) GetPassCreativeTasks(bookId int, studentId, questionId int, isCreative bool) (models.PassCreativeTask, error) {
	return g.CreativeTaskRepo.GetCurrentStudentPass(bookId, uint(studentId), questionId, isCreative)
}

func (g *CreativeTaskService) CreateCreativeTask(question models.CreativeTask, isCreative bool) (int, error) {
	if len(question.Question) > 200 {
		return 0, errors.New("the question's number of letter is overlimitted than 200")
	}
	isKaz := true
	if question.Category != "2 год" && question.Category != "3 год" && question.Category != "4 год" && question.Category != "5 год" &&
		question.Category != "1 класс" && question.Category != "2 класс" && question.Category != "3 класс" && question.Category != "4 класс" {
		isKaz = false

	} else if question.Category != "2 жас" &&
		question.Category != "3 жас" && question.Category != "4 жас" && question.Category != "5 жас" && question.Category != "1 сынып" &&
		question.Category != "2 сынып" && question.Category != "3 сынып" && question.Category != "4 сынып" {
		isKaz = true
	} else {
		return 0, errors.New("not correct category")
	}
	id, err := g.CreativeTaskRepo.CreateCreativeTask(question, isCreative)
	if err != nil {
		return 0, err
	}

	audioLink, err := tts.TextToSpeech(question.Question, "./static/openQuestions/", id, isKaz)
	if err != nil {
		return 0, err
	}

	if err := g.CreativeTaskRepo.AddAudioLinkQuestion(g.imageUrl+audioLink[1:], uint(id), isCreative); err != nil {
		return 0, err
	}

	return id, nil
}

func (g *CreativeTaskService) DeleteCreativeTask(id int, isCreative bool) error {
	return g.CreativeTaskRepo.DeleteCreativeTask(id, isCreative)
}

func (g *CreativeTaskService) UpdateCreativeTask(question interface{}, isCreative bool) error {
	input, ok := question.(map[string]interface{})

	if !ok {
		log.Println("1")
		return errors.New("cannot decode input in update creative task")
	}
	log.Println(input)
	text, ok := input["question"].(string)
	if !ok {
		log.Println("3")
		return errors.New("cannot decode input in update creative task")
	}
	id, ok := input["id"].(float64)
	if !ok {
		return errors.New("cannot decode input in update creative task")
	}

	isKaz := true
	if input["category"] != nil {
		if input["category"] != "2 год" && input["category"] != "3 год" && input["category"] != "4 год" && input["category"] != "5 год" &&
			input["category"] != "1 класс" && input["category"] != "2 класс" && input["category"] != "3 класс" && input["category"] != "4 класс" {
			isKaz = false

		} else if input["category"] != "2 жас" &&
			input["category"] != "3 жас" && input["category"] != "4 жас" && input["category"] != "5 жас" && input["category"] != "1 сынып" &&
			input["category"] != "2 сынып" && input["category"] != "3 сынып" && input["category"] != "4 сынып" {
			isKaz = true
		} else {
			return errors.New("not correct category")
		}
	}
	if err := g.CreativeTaskRepo.UpdateCreativeTask(question, isCreative); err != nil {
		return err
	}

	audioLink, err := tts.TextToSpeech(text, "./static/openQuestions/", int(id), isKaz)
	if err != nil {
		return err
	}

	if err := g.CreativeTaskRepo.AddAudioLinkQuestion(g.imageUrl+audioLink[1:], uint(id), isCreative); err != nil {
		return err
	}
	return nil
}

func (g *CreativeTaskService) GetPassedStudents(ctx context.Context, teacherId int, bookId int, isCreative bool) (models.FullTeacher, error) {
	var fullTeacher models.FullTeacher

	teacher, err := g.teacherRepo.GetById(ctx, uint(teacherId))
	if err != nil {
		return models.FullTeacher{}, fmt.Errorf("CreativeTask service: get all for teacher: %w", err)
	}
	fullTeacher.Teacher = teacher

	classes, err := g.classRepo.GetAll(context.WithValue(ctx, models.TeacherId, teacher.Id))
	if err != nil {
		return models.FullTeacher{}, fmt.Errorf("CreativeTask service: get all for teacher: %w", err)
	}
	fullTeacher.Classes = make([]models.FullClass, len(classes))

	for i, class := range classes {
		fullTeacher.Classes[i].Class = class

		students, err := g.CreativeTaskRepo.GetPassedStudents(int(class.Id), bookId, isCreative)
		if err != nil {
			return models.FullTeacher{}, fmt.Errorf("teacher service: get all for teacher: %w", err)
		}

		fullTeacher.Classes[i].Students = students
	}
	return fullTeacher, nil
}

func (g *CreativeTaskService) GetStudentAllPasses(bookId int, studentId uint, isCreative bool) ([]models.PassCreativeTask, error) {
	return g.CreativeTaskRepo.GetStudentAllPasses(bookId, studentId, isCreative)
}

func (g *CreativeTaskService) PostCommentStudent(ctx context.Context, comment models.CheckCreativePass, isCreative bool) error {
	if len(comment.Comment) > 1000 {
		return errors.New("number letter of comment are overlimmited than 1000")
	}

	if err := g.studentRepo.GivePoints(ctx, comment.StudentId, comment.Point); err != nil {
		return fmt.Errorf("create comment service: give points: %w", err)
	}

	answer, err := g.CreativeTaskRepo.GetAnswerById(comment.AnswerId, isCreative)
	if err != nil {
		return fmt.Errorf("create comment service: get answer of creative task: %w", err)
	}
	book, err := g.bookRepo.Get(context.WithValue(ctx, models.BookId, uint(answer.BookId)))
	if err != nil {
		return fmt.Errorf("create comment service: get book language: %w", err)
	}
	id, err := g.CreativeTaskRepo.PostCommentStudent(comment, isCreative)
	if err != nil {
		return err
	}
	lang := book[0].Language

	isKaz := true
	if lang == "RUS" {
		isKaz = false
	}

	audioLink, err := tts.TextToSpeech(comment.Comment, fmt.Sprintf("./static/comments/%d/", id), 1, isKaz)
	if err != nil {
		return err
	}

	if err := g.CreativeTaskRepo.AddAudioLinkComment(g.imageUrl+audioLink[1:], id, isCreative); err != nil {
		return err
	}
	return nil
}

func (g *CreativeTaskService) GetComments(studentId uint, isCreative bool) ([]models.CreativeNotifications, error) {
	return g.CreativeTaskRepo.GetComments(studentId, g.imageUrl, isCreative)
}
