package http

import (
	"book-school/internal/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) createOpenQuestions(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.AdminRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}
	var (
		question models.CreativeTask
		err      error
	)
	if err = json.NewDecoder(r.Body).Decode(&question); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.CreativeTask.CreateCreativeTask(question, false)
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	if err = json.NewEncoder(w).Encode(id); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) deleteOpenQuestions(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.AdminRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}
	questionId, err := strconv.Atoi(mux.Vars(r)["question_id"])
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
	}
	if questionId < 1 {
		h.errorHandler(w, r, http.StatusBadRequest, "not correct id")
		return
	}
	err = h.service.CreativeTask.DeleteCreativeTask(questionId, false)
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
	}
}

func (h *Handler) updateOpenQuestions(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.AdminRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}
	var (
		question interface{}
		err      error
	)
	if err = json.NewDecoder(r.Body).Decode(&question); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}
	if err = h.service.CreativeTask.UpdateCreativeTask(question, false); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	if err = json.NewEncoder(w).Encode(statusResponse{Status: "OK"}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) getOpenQuestions(w http.ResponseWriter, r *http.Request) {
	category := mux.Vars(r)["category"]
	user := r.Context().Value(models.UserCtx).(models.User)
	questions, err := h.service.CreativeTask.GetCreativeTask(user, category, false)
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(questions); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) completeOpenQuestions(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)
	if user.Role != models.StudentRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	var (
		Answers models.PassCreativeTask
		err     error
	)

	if err = r.ParseMultipartForm(10 << 20); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	jsonDoc := r.FormValue("document")
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}
	if err = json.Unmarshal([]byte(jsonDoc), &Answers); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	Answers.StudentId = user.Id
	if err = h.service.CreativeTask.CreatePassCreativeTask(Answers, false); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(statusResponse{Status: "OK"}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) getAllPassOpenQuestions(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.TeacherRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	studentId, err := strconv.Atoi(mux.Vars(r)["student_id"])
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if studentId < 1 {
		h.errorHandler(w, r, http.StatusBadRequest, "not correct id")
		return
	}

	bookId, err := strconv.Atoi(mux.Vars(r)["book_id"])
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if bookId < 1 {
		h.errorHandler(w, r, http.StatusBadRequest, "not correct id")
		return
	}

	questionId, err := strconv.Atoi(mux.Vars(r)["question_id"])
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if questionId < 1 {
		h.errorHandler(w, r, http.StatusBadRequest, "not correct id")
		return
	}

	answer, err := h.service.CreativeTask.GetPassCreativeTasks(bookId, studentId, questionId, false)
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(answer); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) getOpenPassedStudents(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.TeacherRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}
	bookId, err := strconv.Atoi(mux.Vars(r)["book_id"])
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
	}

	classes, err := h.service.CreativeTask.GetPassedStudents(r.Context(), int(user.Id), bookId, false)
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(classes); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) GetOpenStudentAllPasses(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.TeacherRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	studentId, err := strconv.Atoi(mux.Vars(r)["student_id"])
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if studentId < 1 {
		h.errorHandler(w, r, http.StatusBadRequest, "not correct id")
		return
	}

	bookId, err := strconv.Atoi(mux.Vars(r)["book_id"])
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if bookId < 1 {
		h.errorHandler(w, r, http.StatusBadRequest, "not correct id")
		return
	}

	answers, err := h.service.CreativeTask.GetStudentAllPasses(bookId, uint(studentId), false)
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(answers); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) PostOpenCommentStudent(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.TeacherRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}
	var (
		comment models.CheckCreativePass
		err     error
	)
	if err = json.NewDecoder(r.Body).Decode(&comment); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}
	comment.TeacherId = user.Id
	if err = h.service.CreativeTask.PostCommentStudent(r.Context(), comment, false); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}
	if err := json.NewEncoder(w).Encode(statusResponse{Status: "OK"}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) GetOpenComments(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.StudentRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}
	comments, err := h.service.CreativeTask.GetComments(user.Id, false)
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	if err = json.NewEncoder(w).Encode(comments); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}
