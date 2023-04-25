package http

import (
	"book-school/internal/models"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// @Summary rePassTest
// @Security ApiKeyAuth
// @Tags test
// @Accept json
// @Produce json
// @Param complete_test_id body integer true "complete test id"
// @Success 200 {object} object
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/test/info/{book_id} [get]
func (h *Handler) rePassTest(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.StudentRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	var req struct {
		CompleteTestId int `json:"complete_test_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Test.RePass(r.Context(), user, req.CompleteTestId)
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(statusResponse{"OK"}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary getTestForTeacher
// @Security ApiKeyAuth
// @Tags test
// @Produce json
// @Param book_id path integer true "test id"
// @Success 200 {object} object
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/test/info/{book_id} [get]
func (h *Handler) getTestForTeacher(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.TeacherRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	bookId, err := strconv.Atoi(mux.Vars(r)["book_id"])
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	tests, err := h.service.Test.GetTestForTeacher(r.Context(), uint(bookId), user.Id)
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(tests); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary getComleteTestForStudent
// @Security ApiKeyAuth
// @Tags test
// @Produce json
// @Param test_id path integer true "test id"
// @Param student_id path integer true "student id"
// @Success 200 {object} models.CompleteTest
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/test/info/{test_id}/{student_id} [get]
func (h *Handler) getComleteTestForStudent(w http.ResponseWriter, r *http.Request) {
	testId, err := strconv.Atoi(mux.Vars(r)["test_id"])
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	studentId, err := strconv.Atoi(mux.Vars(r)["student_id"])
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	test, err := h.service.Test.GetCompleteTest(r.Context(), uint(testId), uint(studentId))
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
			return
		}
	}

	if err = json.NewEncoder(w).Encode(test); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) createNewTest(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.AdminRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	var err error
	var test models.Test

	if err = r.ParseMultipartForm(10 << 20); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	jsonDoc, ok := r.Form["document"]
	if !ok {
		h.errorHandler(w, r, http.StatusBadRequest, "document field not found")
		return
	}
	questions := r.MultipartForm.File["question"]
	answers := r.MultipartForm.File["answer"]

	if err = json.Unmarshal([]byte(jsonDoc[0]), &test); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err = h.service.Test.Create(r.Context(), test, questions, answers); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(statusResponse{Status: "OK"}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary getTestByBookId
// @Security ApiKeyAuth
// @Tags test
// @Produce json
// @Param book_id path integer true "book id"
// @Success 200 {object} models.Test
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/test/{book_id} [get]
func (h *Handler) getTestByBookId(w http.ResponseWriter, r *http.Request) {
	bookId, err := strconv.Atoi(mux.Vars(r)["book_id"])
	if err != nil {
		h.errorHandler(w, r, http.StatusNotFound, err.Error())
		return
	}

	ctx := context.WithValue(r.Context(), models.BookId, uint(bookId))
	tests, err := h.service.Test.Get(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			h.errorHandler(w, r, http.StatusBadRequest, err.Error())
			return
		}
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(tests); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary completeTest
// @Security ApiKeyAuth
// @Tags test
// @Accept json
// @Produce json
// @Param test body models.CompleteTest true "test info"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/test/complete [post]
func (h *Handler) completeTest(w http.ResponseWriter, r *http.Request) {
	var err error

	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.StudentRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	var test models.CompleteTest
	if err = json.NewDecoder(r.Body).Decode(&test); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}
	test.StudentId = user.Id

	resp, err := h.service.Test.CompleteTest(r.Context(), test)
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(resp); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary deleteTest
// @Security ApiKeyAuth
// @Tags test
// @Produce json
// @Param book_id path integer true "book id"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/test/delete/{book_id} [delete]
func (h *Handler) deleteTest(w http.ResponseWriter, r *http.Request) {
	testId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.AdminRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	if err := h.service.Test.Delete(r.Context(), uint(testId)); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(statusResponse{Status: "OK"}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}
