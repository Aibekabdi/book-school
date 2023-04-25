package http

import (
	"book-school/internal/service"
	"net/http"

	_ "book-school/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.Use(h.loggingMiddleWare)
	router.Use(h.corsMiddleWare)

	fs := http.FileServer(http.Dir("./static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		fs.ServeHTTP(w, r)
	})))

	router.PathPrefix("/swagger").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	auth := router.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/sign-in", h.signIn).Methods(http.MethodPost, http.MethodOptions)
	auth.HandleFunc("/admin/sign-in", h.adminSignIn).Methods(http.MethodPost, http.MethodOptions)

	schoolAuth := auth.PathPrefix("/school").Subrouter()
	schoolAuth.HandleFunc("/sign-up", h.schoolSignUp).Methods(http.MethodPost, http.MethodOptions)

	teacherAuth := auth.PathPrefix("/teacher").Subrouter()
	teacherAuth.HandleFunc("/sign-up", h.teacherSignUp).Methods(http.MethodPost, http.MethodOptions)

	api := router.PathPrefix("/api").Subrouter()
	api.Use(h.userIdentity)

	school := api.PathPrefix("/school").Subrouter()
	school.HandleFunc("/profile", h.getSchoolInfo).Methods(http.MethodGet, http.MethodOptions)
	school.HandleFunc("/profile", h.updateSchoolInfo).Methods(http.MethodPatch, http.MethodOptions)
	school.HandleFunc("/all", h.getAllSchools).Methods(http.MethodGet, http.MethodOptions)
	school.HandleFunc("/delete/{id}", h.deleteSchool).Methods(http.MethodDelete, http.MethodOptions)

	teacher := api.PathPrefix("/teacher").Subrouter()
	teacher.HandleFunc("/profile", h.getTeacherInfo).Methods(http.MethodGet, http.MethodOptions)
	teacher.HandleFunc("/profile", h.updateTeacherInfo).Methods(http.MethodPatch, http.MethodOptions)
	teacher.HandleFunc("/all", h.getAllTeachers).Methods(http.MethodGet, http.MethodOptions)
	teacher.HandleFunc("/create", h.createNewTeacher).Methods(http.MethodPost, http.MethodOptions)
	teacher.HandleFunc("/delete/{id}", h.deleteTeacher).Methods(http.MethodDelete, http.MethodOptions)

	class := api.PathPrefix("/class").Subrouter()
	class.HandleFunc("/create", h.createNewClass).Methods(http.MethodPost, http.MethodOptions)
	class.HandleFunc("/all", h.getAllClasses).Methods(http.MethodGet, http.MethodOptions)
	class.HandleFunc("/delete/{id}", h.deleteClass).Methods(http.MethodDelete, http.MethodOptions)
	class.HandleFunc("/stats", h.getClassStats).Methods(http.MethodGet, http.MethodOptions)
	class.HandleFunc("/stats/total", h.getStatsFromAllClass).Methods(http.MethodGet, http.MethodOptions)

	student := api.PathPrefix("/student").Subrouter()
	student.HandleFunc("/create", h.createNewStudent).Methods(http.MethodPost, http.MethodOptions)
	student.HandleFunc("/profile", h.getStudentInfo).Methods(http.MethodGet, http.MethodOptions)
	student.HandleFunc("/profile", h.updateStudentInfo).Methods(http.MethodPatch, http.MethodOptions)
	student.HandleFunc("/delete/{id}", h.deleteStudent).Methods(http.MethodDelete, http.MethodOptions)
	student.HandleFunc("/stats", h.studentStats).Methods(http.MethodGet, http.MethodOptions)

	body := student.PathPrefix("/body").Subrouter()
	body.HandleFunc("/current", h.getCurrentBudy).Methods(http.MethodGet, http.MethodOptions)
	body.HandleFunc("/all", h.getBuyedBodyForStudent).Methods(http.MethodGet, http.MethodOptions)
	body.HandleFunc("/update/{from_id}/{to_id}", h.updateCurrentBody).Methods(http.MethodPatch, http.MethodOptions)

	// admin := api.PathPrefix("/admin").Subrouter()
	// admin.HandleFunc("")

	shop := api.PathPrefix("/shop").Subrouter()
	shop.HandleFunc("/all", h.getAllBody).Methods(http.MethodGet, http.MethodOptions)
	shop.HandleFunc("/buy/{body_id}", h.buyBody).Methods(http.MethodPost, http.MethodOptions)
	shop.HandleFunc("/create", h.createNewBody).Methods(http.MethodPost, http.MethodOptions)

	test := api.PathPrefix("/test").Subrouter()
	test.HandleFunc("/{book_id}", h.getTestByBookId).Methods(http.MethodGet, http.MethodOptions)
	test.HandleFunc("/complete", h.completeTest).Methods(http.MethodPost, http.MethodOptions)
	test.HandleFunc("/repass", h.rePassTest).Methods(http.MethodPost, http.MethodOptions)
	test.HandleFunc("/create", h.createNewTest).Methods(http.MethodPost, http.MethodOptions)
	test.HandleFunc("/delete/{id}", h.deleteTest).Methods(http.MethodDelete, http.MethodOptions)
	test.HandleFunc("/info/{test_id}/{student_id}", h.getComleteTestForStudent).Methods(http.MethodGet, http.MethodOptions)
	test.HandleFunc("/info/{book_id}", h.getTestForTeacher).Methods(http.MethodGet, http.MethodOptions)

	books := api.PathPrefix("/books").Subrouter()
	books.HandleFunc("", h.getAllBooksForTest).Methods(http.MethodGet, http.MethodOptions)
	books.HandleFunc("/total", h.totalPages).Methods(http.MethodGet, http.MethodOptions)
	books.HandleFunc("/delete/{book_id}", h.deleteBook).Methods(http.MethodDelete, http.MethodOptions)
	books.HandleFunc("/create", h.createBook).Methods(http.MethodPost, http.MethodOptions)
	books.HandleFunc("/{id}", h.getBookById).Methods(http.MethodGet, http.MethodOptions)
	books.HandleFunc("/complete/{id}", h.completeBook).Methods(http.MethodPost, http.MethodOptions)
	all := books.PathPrefix("/all").Subrouter()
	all.HandleFunc("/all", h.getAllBooks).Methods(http.MethodGet, http.MethodOptions)

	audio := api.PathPrefix("/audio").Subrouter()
	audio.HandleFunc("/complete/{book_id}", h.completeAudio).Methods(http.MethodPost, http.MethodOptions)

	creative := api.PathPrefix("/creative").Subrouter()
	CreativeTask := creative.PathPrefix("/tasks").Subrouter()
	CreativeTask.HandleFunc("/get", h.getCreativeTask).Methods(http.MethodGet, http.MethodOptions)
	CreativeTask.HandleFunc("/get/{category}", h.getCreativeTask).Methods(http.MethodGet, http.MethodOptions)
	CreativeTask.HandleFunc("/create", h.createCreativeTask).Methods(http.MethodPost, http.MethodOptions)
	CreativeTask.HandleFunc("/delete/{question_id}", h.deleteCreativeTask).Methods(http.MethodDelete, http.MethodOptions)
	CreativeTask.HandleFunc("/update", h.updateCreativeTask).Methods(http.MethodPatch, http.MethodOptions)

	PassCreativeTask := creative.PathPrefix("/pass").Subrouter()
	PassCreativeTask.HandleFunc("/create", h.completeCreativeTask).Methods(http.MethodPost, http.MethodOptions)
	PassCreativeTask.HandleFunc("/get/{book_id}/{student_id}/{question_id}", h.getAllPassCreativeTask).Methods(http.MethodGet, http.MethodOptions)

	CheckTasks := creative.PathPrefix("/check").Subrouter()
	CheckTasks.HandleFunc("/get/all/{book_id}", h.getCreativePassedStudents).Methods(http.MethodGet, http.MethodOptions)
	CheckTasks.HandleFunc("/get/student/passes/{student_id}/{book_id}", h.GetCreativeStudentAllPasses).Methods(http.MethodGet, http.MethodOptions)
	CheckTasks.HandleFunc("/comment", h.PostCreativeCommentStudent).Methods(http.MethodPost, http.MethodOptions)
	CheckTasks.HandleFunc("/comment", h.GetCreativeComments).Methods(http.MethodGet, http.MethodOptions)

	open := api.PathPrefix("/open").Subrouter()
	OpenQuestions := open.PathPrefix("/tasks").Subrouter()
	OpenQuestions.HandleFunc("/get", h.getOpenQuestions).Methods(http.MethodGet, http.MethodOptions)
	OpenQuestions.HandleFunc("/get/{category}", h.getOpenQuestions).Methods(http.MethodGet, http.MethodOptions)
	OpenQuestions.HandleFunc("/create", h.createOpenQuestions).Methods(http.MethodPost, http.MethodOptions)
	OpenQuestions.HandleFunc("/delete/{question_id}", h.deleteOpenQuestions).Methods(http.MethodDelete, http.MethodOptions)
	OpenQuestions.HandleFunc("/update", h.updateOpenQuestions).Methods(http.MethodPatch, http.MethodOptions)

	PassOpenQuestions := open.PathPrefix("/pass").Subrouter()
	PassOpenQuestions.HandleFunc("/create", h.completeOpenQuestions).Methods(http.MethodPost, http.MethodOptions)
	PassOpenQuestions.HandleFunc("/get/{book_id}/{student_id}/{question_id}", h.getAllPassOpenQuestions).Methods(http.MethodGet, http.MethodOptions)

	Open := open.PathPrefix("/check").Subrouter()
	Open.HandleFunc("/get/all/{book_id}", h.getOpenPassedStudents).Methods(http.MethodGet, http.MethodOptions)
	Open.HandleFunc("/get/student/passes/{student_id}/{book_id}", h.GetOpenStudentAllPasses).Methods(http.MethodGet, http.MethodOptions)
	Open.HandleFunc("/comment", h.PostOpenCommentStudent).Methods(http.MethodPost, http.MethodOptions)
	Open.HandleFunc("/comment", h.GetOpenComments).Methods(http.MethodGet, http.MethodOptions)
	return router
}
