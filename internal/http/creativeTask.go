package http

import (
	"book-school/internal/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// @Summary create update delete createCreativeTask
// @Security ApiKeyAuth
// @Tags CreativeTaskRETURNING
// @Accept json
// @Produce json
// @Param CreativeTask body models.CreativeTask true "open questions info"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/creative/tasks/create [post]
func (h *Handler) createCreativeTask(w http.ResponseWriter, r *http.Request) {
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
	id, err := h.service.CreativeTask.CreateCreativeTask(question, true)
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	if err = json.NewEncoder(w).Encode(id); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary create update delete deletePassCreativeTask
// @Security ApiKeyAuth
// @Tags CreativeTask
// @Accept json
// @Produce json
// @Param PassCreativeTask body models.PassCreativeTask true "open answers info"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/creative/pass/delete/{question_id} [delete]
func (h *Handler) deleteCreativeTask(w http.ResponseWriter, r *http.Request) {
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
	err = h.service.CreativeTask.DeleteCreativeTask(questionId, true)
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
	}
}

func (h *Handler) updateCreativeTask(w http.ResponseWriter, r *http.Request) {
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
	if err = h.service.CreativeTask.UpdateCreativeTask(question, true); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	if err = json.NewEncoder(w).Encode(statusResponse{Status: "OK"}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary getCreativeTask
// @Security ApiKeyAuth
// @Tags CreativeTask
// @Accept json
// @Produce json
// @Success 200 {object} []models.CreativeTask
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/creative/tasks/get [get]
func (h *Handler) getCreativeTask(w http.ResponseWriter, r *http.Request) {
	category := mux.Vars(r)["category"]
	user := r.Context().Value(models.UserCtx).(models.User)
	questions, err := h.service.CreativeTask.GetCreativeTask(user, category, true)
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(questions); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary create update delete createPassCreativeTask
// @Security ApiKeyAuth
// @Tags CreativeTask
// @Accept json
// @Produce json
// @Param PassCreativeTask body models.PassCreativeTask true "open answers info"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/creative/pass/create [post]
func (h *Handler) completeCreativeTask(w http.ResponseWriter, r *http.Request) {
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
	file, _, err := r.FormFile("art")
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}
	Answers.Img = file
	Answers.StudentId = user.Id
	if err = h.service.CreativeTask.CreatePassCreativeTask(Answers, true); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(statusResponse{Status: "OK"}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary GetPassCreativeTasks
// @Security ApiKeyAuth
// @Tags CreativeTask
// @Accept json
// @Produce json
// @Success 200 {object} models.PassCreativeTask
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/creative/pass/get/{student_id}/{book_id} [get]
func (h *Handler) getAllPassCreativeTask(w http.ResponseWriter, r *http.Request) {
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

	answer, err := h.service.CreativeTask.GetPassCreativeTasks(bookId, studentId, questionId, true)
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(answer); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) getCreativePassedStudents(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.TeacherRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}
	bookId, err := strconv.Atoi(mux.Vars(r)["book_id"])
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
	}

	classes, err := h.service.CreativeTask.GetPassedStudents(r.Context(), int(user.Id), bookId, true)
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(classes); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) GetCreativeStudentAllPasses(w http.ResponseWriter, r *http.Request) {
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

	answers, err := h.service.CreativeTask.GetStudentAllPasses(bookId, uint(studentId), true)
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(answers); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) PostCreativeCommentStudent(w http.ResponseWriter, r *http.Request) {
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
	if err = h.service.CreativeTask.PostCommentStudent(r.Context(), comment, true); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}
	if err := json.NewEncoder(w).Encode(statusResponse{Status: "OK"}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) GetCreativeComments(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.StudentRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}
	comments, err := h.service.CreativeTask.GetComments(user.Id, true)
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	if err = json.NewEncoder(w).Encode(comments); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}
