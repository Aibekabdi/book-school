package service

import (
	"book-school/internal/models"
	"book-school/internal/repository"
	"book-school/pkg/config"
	"context"
	"mime/multipart"
)

type Admin interface {
	Create(ctx context.Context, admin models.Admin) error
	GetById(ctx context.Context, id uint) (models.Admin, error)
}

type Auth interface {
	SignIn(ctx context.Context, name, password, role string) (string, error)
	ParseToken(token string) (models.User, error)
}

type School interface {
	Create(ctx context.Context, school models.School) error
	Update(ctx context.Context, user models.User, update models.SchoolUpdate) error
	GetById(ctx context.Context, schoolId uint) (models.School, error)
	GetAllForSchool(ctx context.Context, schoolId uint) (models.FullSchool, error)
	GetAllForAdmin(ctx context.Context) ([]models.FullSchool, error)
	Delete(ctx context.Context, schoolId uint) error
}

type Teacher interface {
	Create(ctx context.Context, teacher models.Teacher, private bool, schoolId uint) error
	GetById(ctx context.Context, id uint) (models.Teacher, error)
	Delete(ctx context.Context, teacherId, schoolId uint) error
	Update(ctx context.Context, user models.User, update models.TeacherUpdate) error
	GetAllForTeacher(ctx context.Context, teacherId uint) (models.FullTeacher, error)
}

type Class interface {
	Create(ctx context.Context, class models.Class, user models.User) error
	GetStats(ctx context.Context, user models.User) ([]models.ClassStats, error)
	Delete(ctx context.Context, classId uint, user models.User) error
	GetStatsTotal(ctx context.Context, user models.User) ([]models.Stats, error)
}

type Student interface {
	Create(ctx context.Context, student models.Student, user models.User) error
	GetById(ctx context.Context, id uint) (models.Student, error)
	Delete(ctx context.Context, studentId, teacherId uint) error
	Update(ctx context.Context, user models.User, update models.StudentUpdate) error
	GetStats(ctx context.Context, id uint) (models.Stats, error)
}

type Book interface {
	GetAllForTest(ctx context.Context) ([]models.Book, error)
	GetAll(ctx context.Context, user models.User) ([]models.BooksStruct, error)
	GetJsonBook(id int) (models.Book, error)
	CreateBook(input *models.Book, file multipart.File) error
	DeleteBook(id int) error
	Complete(ctx context.Context, bookId, studentId uint) (uint, error)
	TotalPages(ctx context.Context, user models.User) (int, error)
}

type Audio interface {
	Complete(ctx context.Context, bookId, studentId uint) (uint, error)
}

type Test interface {
	RePass(ctx context.Context, user models.User, testId int) error
	Create(ctx context.Context, test models.Test, questions, answers []*multipart.FileHeader) error
	Get(ctx context.Context) ([]models.Test, error)
	CompleteTest(ctx context.Context, test models.CompleteTest) (models.CompleteTestResp, error)
	Delete(ctx context.Context, testId uint) error
	GetTestForTeacher(ctx context.Context, testId, teacherId uint) ([]models.TestForTeacher, error)
	GetCompleteTest(ctx context.Context, testId, studentId uint) (models.CompleteTest, error)
}

type CreativeTask interface {
	GetCreativeTask(user models.User, category string, isCreative bool) ([]models.CreativeTask, error)
	CreateCreativeTask(question models.CreativeTask, isCreative bool) (int, error)
	DeleteCreativeTask(id int, isCreative bool) error
	UpdateCreativeTask(question interface{}, isCreative bool) error
	CreatePassCreativeTask(answer models.PassCreativeTask, isCreative bool) error
	GetPassCreativeTasks(bookId int, studentId, questionId int, isCreative bool) (models.PassCreativeTask, error)
	GetPassedStudents(ctx context.Context, teacherId int, bookId int, isCreative bool) (models.FullTeacher, error)
	GetStudentAllPasses(bookId int, studentId uint, isCreative bool) ([]models.PassCreativeTask, error)
	PostCommentStudent(ctx context.Context, comment models.CheckCreativePass, isCreative bool) error
	GetComments(studentId uint, isCreative bool) ([]models.CreativeNotifications, error)
}

type Shop interface {
	Create(ctx context.Context, image, imageIcon *multipart.FileHeader, price, name, part string) (models.Body, error)
	GetAll(ctx context.Context, studentId uint) (models.ShopInfo, error)
	Buy(ctx context.Context, studentId uint, bodyId uint) (uint, error)
	GetAllBuyed(ctx context.Context, studentId uint) (models.ShopInfo, error)
	GetCurrentBody(ctx context.Context, studentId uint) (models.ShopInfo, error)
	UpdateCurrentBody(ctx context.Context, studentId, fromId, toId uint) error
}

type Service struct {
	Auth         Auth
	School       School
	Class        Class
	Teacher      Teacher
	Student      Student
	Book         Book
	Audio        Audio
	Admin        Admin
	Test         Test
	CreativeTask CreativeTask
	Shop         Shop
}

func NewService(r *repository.Repository, cfg *config.Conf) *Service {
	imageUrl := "http://" + cfg.Api.Host + ":" + cfg.Api.Port
	return &Service{
		Auth:         newAuthService(r.School, r.Teacher, r.Student, r.Admin),
		School:       newSchoolService(r.School, r.Teacher, r.Class, r.Student),
		Class:        newClassService(r.Class, r.Student, r.Teacher, r.School),
		Teacher:      newTeacherService(r.Teacher, r.Class, r.Student, r.School),
		Student:      newStudentService(r.Student, r.Class, r.Shop),
		Book:         newBookService(r.Book, r.Student, r.Class, imageUrl),
		Audio:        newAudioService(r.Audio, r.Student),
		Admin:        newAdminService(r.Admin),
		Test:         newTestService(r.Test, r.Question, r.Answer, r.Student, r.Class, r.Book, imageUrl),
		CreativeTask: newCreativeTaskService(r.CreativeTask, r.Student, r.Class, r.Teacher, r.School, r.Book, imageUrl),
		Shop:         newShopService(r.Shop, r.Student, cfg),
	}
}
