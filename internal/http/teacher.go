package http

import (
	"book-school/internal/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// @Summary deleteTeacher
// @Security ApiKeyAuth
// @Tags teacher
// @Produce json
// @Param id path int true "teacher id"
// @Success 200 {object} object
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/teacher/delete/{id} [delete]
func (h *Handler) deleteTeacher(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.SchoolRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	teacherId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		h.errorHandler(w, r, http.StatusNotFound, err.Error())
		return
	}

	err = h.service.Teacher.Delete(r.Context(), uint(teacherId), user.Id)
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(statusResponse{Status: "OK"}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary teacherSignUp
// @Tags teacher
// @Accept json
// @Produce json
// @Param newTeacher body models.Teacher true "teacher info"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/teacher/sign-up [post]
func (h *Handler) teacherSignUp(w http.ResponseWriter, r *http.Request) {
	var (
		newTeacher models.Teacher
		err        error
	)
	if err = json.NewDecoder(r.Body).Decode(&newTeacher); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.service.Teacher.Create(r.Context(), newTeacher, true, 1); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(statusResponse{Status: "OK"}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary createNewTeacher
// @Security ApiKeyAuth
// @Tags teacher
// @Accept json
// @Produce json
// @Param newTeacher body models.Teacher true "teacher info"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/teacher/create [post]
func (h *Handler) createNewTeacher(w http.ResponseWriter, r *http.Request) {
	var (
		newTeacher models.Teacher
		err        error
	)

	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.SchoolRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&newTeacher); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.service.Teacher.Create(r.Context(), newTeacher, false, user.Id); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(statusResponse{Status: "OK"}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary getAllTeachers
// @Security ApiKeyAuth
// @Tags teacher
// @Produce json
// @Success 200 {object} models.FullSchool
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/teacher/all [get]
func (h *Handler) getAllTeachers(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.SchoolRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	school, err := h.service.School.GetAllForSchool(r.Context(), user.Id)
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(school); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary getTeacherInfo
// @Security ApiKeyAuth
// @Tags teacher
// @Produce json
// @Success 200 {object} models.Teacher
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/teacher/profile [get]
func (h *Handler) getTeacherInfo(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.TeacherRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	teacher, err := h.service.Teacher.GetById(r.Context(), user.Id)
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(teacher); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary updateTeacherInfo
// @Security ApiKeyAuth
// @Tags teacher
// @Accept json
// @Produce json
// @Param update body models.TeacherUpdate true "teacher update info"
// @Success 200 {object} models.Teacher
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/teacher/profile [patch]
func (h *Handler) updateTeacherInfo(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.TeacherRole && user.Role != models.SchoolRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	var update models.TeacherUpdate
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Teacher.Update(r.Context(), user, update); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(statusResponse{Status: "OK"}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}
