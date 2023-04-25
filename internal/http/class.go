package http

import (
	"book-school/internal/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// @Summary deleteClass
// @Security ApiKeyAuth
// @Tags class
// @Produce json
// @Param id path int true "class id"
// @Success 200 {object} object
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/class/delete/{id} [delete]
func (h *Handler) deleteClass(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.SchoolRole || user.Role != models.TeacherRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	classId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		h.errorHandler(w, r, http.StatusNotFound, err.Error())
		return
	}

	err = h.service.Class.Delete(r.Context(), uint(classId), user)
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(statusResponse{Status: "OK"}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary createNewClass
// @Security ApiKeyAuth
// @Tags class
// @Accept json
// @Produce json
// @Param newClass body models.Class true "class info"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/class/create [post]
func (h *Handler) createNewClass(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.SchoolRole && user.Role != models.TeacherRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}
	var (
		newClass models.Class
		err      error
	)

	if err = json.NewDecoder(r.Body).Decode(&newClass); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.service.Class.Create(r.Context(), newClass, user); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(statusResponse{Status: "OK"}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary getAllClasses
// @Security ApiKeyAuth
// @Tags class
// @Produce json
// @Success 200 {object} models.FullTeacher
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/class/all [get]
func (h *Handler) getAllClasses(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.TeacherRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	teacher, err := h.service.Teacher.GetAllForTeacher(r.Context(), user.Id)
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(teacher); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary getClassStats
// @Security ApiKeyAuth
// @Tags class
// @Produce json
// @Success 200 {object} []models.ClassStats
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/class/stats [get]
func (h *Handler) getClassStats(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.TeacherRole && user.Role != models.SchoolRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	stats, err := h.service.Class.GetStats(r.Context(), user)
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(stats); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary getStatsFromAllClass
// @Security ApiKeyAuth
// @Tags class
// @Produce json
// @Success 200 {object} []models.Stats
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/class/stats/total [get]
func (h *Handler) getStatsFromAllClass(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.TeacherRole && user.Role != models.SchoolRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	stats, err := h.service.Class.GetStatsTotal(r.Context(), user)
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(stats); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}
