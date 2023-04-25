package http

import (
	"book-school/internal/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// @Summary deleteStudent
// @Security ApiKeyAuth
// @Tags student
// @Produce json
// @Param id path int true "student id"
// @Success 200 {object} object
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/student/delete/{id} [delete]
func (h *Handler) deleteStudent(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.TeacherRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	studentId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		h.errorHandler(w, r, http.StatusNotFound, err.Error())
		return
	}

	err = h.service.Student.Delete(r.Context(), uint(studentId), user.Id)
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(statusResponse{Status: "OK"}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary updateCurrentBody
// @Security ApiKeyAuth
// @Tags student
// @Produce json
// @Param from_id path int true "from body id"
// @Param to_id path int true "to body id"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/student/body/update/{from_id}/{to_id} [patch]
func (h *Handler) updateCurrentBody(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.StudentRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	fromId, err := strconv.Atoi(mux.Vars(r)["from_id"])
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	toId, err := strconv.Atoi(mux.Vars(r)["to_id"])
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Shop.UpdateCurrentBody(r.Context(), user.Id, uint(fromId), uint(toId))
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(statusResponse{Status: "OK"}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary getBuyedBodyForStudent
// @Security ApiKeyAuth
// @Tags student
// @Produce json
// @Success 200 {object} models.ShopInfo
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/student/body/current [get]
func (h *Handler) getCurrentBudy(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.StudentRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	info, err := h.service.Shop.GetCurrentBody(r.Context(), user.Id)
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(info); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary getBuyedBodyForStudent
// @Security ApiKeyAuth
// @Tags student
// @Produce json
// @Success 200 {object} models.ShopInfo
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/student/body/all [get]
func (h *Handler) getBuyedBodyForStudent(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.StudentRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	info, err := h.service.Shop.GetAllBuyed(r.Context(), user.Id)
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(info); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary createNewStudent
// @Security ApiKeyAuth
// @Tags student
// @Accept json
// @Produce json
// @Param newStudent body models.Student true "student info"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/student/create [post]
func (h *Handler) createNewStudent(w http.ResponseWriter, r *http.Request) {
	var (
		newStudent models.Student
		err        error
	)

	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.TeacherRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&newStudent); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.service.Student.Create(r.Context(), newStudent, user); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(statusResponse{Status: "OK"}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) studentStats(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.StudentRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	stats, err := h.service.Student.GetStats(r.Context(), user.Id)
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(stats); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary getStudentInfo
// @Security ApiKeyAuth
// @Tags student
// @Produce json
// @Success 200 {object} models.Student
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/student/profile [get]
func (h *Handler) getStudentInfo(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.StudentRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	student, err := h.service.Student.GetById(r.Context(), user.Id)
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(student); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary updateStudentInfo
// @Security ApiKeyAuth
// @Tags student
// @Accept json
// @Produce json
// @Param update body models.StudentUpdate true "student update info"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/student/profile [patch]
func (h *Handler) updateStudentInfo(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.TeacherRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	var update models.StudentUpdate
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Student.Update(r.Context(), user, update); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(statusResponse{Status: "OK"}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}
