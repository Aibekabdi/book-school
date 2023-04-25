package repository

import (
	"book-school/internal/models"
	"context"
	"database/sql"
)

const (
	schoolTable          = "schools"
	classTable           = "classes"
	teacherTable         = "teachers"
	studentTable         = "students"
	adminTable           = "admins"
	testTable            = "tests"
	completeTestsTable   = "complete_tests"
	questionTable        = "questions"
	answerTable          = "answers"
	completeAnswersTable = "complete_answers"
	bookTable            = "books"
	completeBooksTable   = "complete_books"
	completeAudioTable   = "complete_audios"
	bodyTable            = "body"
	currentBodyTable     = "current_body"
	buyedTable           = "buyed"
)

type School interface {
	Create(ctx context.Context, school models.School) error
	GetSchool(ctx context.Context, name string) (uint, string, error)
	Update(ctx context.Context, update models.SchoolUpdate) error
	GetAll(ctx context.Context) ([]models.School, error)
	Delete(ctx context.Context, schoolId uint) error
}

type Class interface {
	Create(ctx context.Context, class models.Class) error
	GetAll(ctx context.Context) ([]models.Class, error)
	GetClass(ctx context.Context, name, grade string, schoolId uint) (models.Class, error)
	Delete(ctx context.Context, classId uint) error
	GetStats(ctx context.Context, classId uint) ([]models.Stats, error)
}

type Teacher interface {
	Create(ctx context.Context, teacher models.Teacher) error
	GetTeacher(ctx context.Context, name string) (uint, string, error)
	GetById(ctx context.Context, id uint) (models.Teacher, error)
	GetAll(ctx context.Context) ([]models.Teacher, error)
	Update(ctx context.Context, update models.TeacherUpdate) error
	Delete(ctx context.Context, teacherId uint) error
}

type Book interface {
	GetAllForTest(ctx context.Context) ([]models.Book, error)
	Get(ctx context.Context) ([]models.Book, error)
	CreateBook(input *models.Book, hashed string) (int, error)
	DeleteBook(hashed string) error
	GetBookHashedId(id int) (string, error)
	Complete(ctx context.Context, bookId, studentId, points uint) error
	CheckCompleteBook(ctx context.Context, bookId, studentId uint) error
	GetAll(ctx context.Context, category []string) ([]models.Book, error)
}

type Audio interface {
	Complete(ctx context.Context, bookId, studentId, points uint) error
	CheckCompleteAudio(ctx context.Context, bookId, studentId uint) error
}

type Student interface {
	Create(ctx context.Context, student models.Student) (uint, error)
	Delete(ctx context.Context, id uint) error
	GetStudent(ctx context.Context, name string) (uint, string, error)
	GetById(ctx context.Context, id uint) (models.Student, error)
	GetAll(ctx context.Context) ([]models.Student, error)
	Update(ctx context.Context, update models.StudentUpdate) error
	GivePoints(ctx context.Context, studentId, points uint) error
	TakePoints(ctx context.Context, studentId, points uint) error
	SetBody(ctx context.Context, studentId, headId, chestId, legsId, armsId uint) error
	UpdateBody(ctx context.Context, studentId, from, to uint, part string) error
}

type Admin interface {
	GetAdmin(ctx context.Context, name string) (uint, string, error)
	GetById(ctx context.Context, id uint) (models.Admin, error)
	Create(ctx context.Context, admin models.Admin) error
}

type Test interface {
	Create(ctx context.Context, bookId uint, lang string) (uint, error)
	Get(ctx context.Context) ([]models.Test, error)
	Delete(ctx context.Context, testId uint) error
	SaveTest(ctx context.Context, test models.CompleteTest) (uint, error)
	DeleteSavedTest(ctx context.Context, id uint) error
	GetCompleteTest(ctx context.Context) (models.CompleteTest, error)
}

type Question interface {
	Create(ctx context.Context, question models.Question) (uint, error)
	GetByTestId(ctx context.Context, testId uint) ([]models.Question, error)
}

type Answer interface {
	Create(ctx context.Context, answer models.Answer) error
	GetByQuestionId(ctx context.Context, questionId uint) ([]models.Answer, error)
	SaveAnswers(ctx context.Context, answer models.CompleteAnswers) error
	GetCompleteAnswers(ctx context.Context, completeTestId uint) ([]models.CompleteAnswers, error)
}

type CreativeTask interface {
	GetCreativeTask(user models.User, category string, isCreative bool) ([]models.CreativeTask, error)
	CreatePassCreativeTask(answers models.PassCreativeTask, isCreative bool) error
	GetCurrentStudentPass(bookId int, studentId uint, questionId int, isCreative bool) (models.PassCreativeTask, error)
	CreateCreativeTask(question models.CreativeTask, isCreative bool) (int, error)
	DeleteCreativeTask(id int, isCreative bool) error
	UpdateCreativeTask(question interface{}, isCreative bool) error
	GetPassedStudents(teacherId int, bookId int, isCreative bool) ([]models.Student, error)
	GetStudentAllPasses(bookId int, studentId uint, isCreative bool) ([]models.PassCreativeTask, error)
	PostCommentStudent(comment models.CheckCreativePass, isCreative bool) (uint, error)
	GetComments(studentId uint, imageUrl string, isCreative bool) ([]models.CreativeNotifications, error)
	GetAnswerById(answerId uint, isCreative bool) (models.PassCreativeTask, error)
	AddAudioLinkComment(audio string, commentId uint, isCreative bool) error
	AddAudioLinkQuestion(audio string, questionId uint, isCreative bool) error
}

type Shop interface {
	Create(ctx context.Context, body models.Body) (uint, error)
	GetAll(ctx context.Context) ([]models.Body, error)
	GetOneBy(ctx context.Context) (models.Body, error)
	Buy(ctx context.Context, studentId, bodyId uint) error
}

type Repository struct {
	School       School
	Class        Class
	Teacher      Teacher
	Student      Student
	Book         Book
	Audio        Audio
	Admin        Admin
	Test         Test
	Question     Question
	Answer       Answer
	CreativeTask CreativeTask
	Shop         Shop
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		School:       newSchoolRepository(db),
		Class:        newClassRepository(db),
		Teacher:      newTeacherRepository(db),
		Student:      newStudentRepository(db),
		Book:         newBookRepository(db),
		Audio:        newAudioRepository(db),
		Admin:        newAdminRepository(db),
		Test:         newTestRepository(db),
		Question:     newQuestionRepository(db),
		Answer:       newAnswerRepository(db),
		CreativeTask: newCreativeTaskRepository(db),
		Shop:         newShopRepository(db),
	}
}
