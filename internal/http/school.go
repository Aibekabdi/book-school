package http

import (
	"book-school/internal/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// @Summary schoolSignUp
// @Tags school
// @Accept json
// @Produce json
// @Param newSchool body models.School true "school sign up info"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/school/sign-up [post]
func (h *Handler) schoolSignUp(w http.ResponseWriter, r *http.Request) {
	var (
		newSchool models.School
		err       error
	)
	if err = json.NewDecoder(r.Body).Decode(&newSchool); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.service.School.Create(r.Context(), newSchool); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(statusResponse{Status: "OK"}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary getSchoolInfo
// @Security ApiKeyAuth
// @Tags school
// @Produce json
// @Success 200 {object} models.School
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/school/profile [get]
func (h *Handler) getSchoolInfo(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.SchoolRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	school, err := h.service.School.GetById(r.Context(), user.Id)
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(school); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary updateSchoolInfo
// @Security ApiKeyAuth
// @Tags school
// @Accept json
// @Produce json
// @Param update body models.SchoolUpdate true "school update info"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/school/profile [patch]
func (h *Handler) updateSchoolInfo(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.SchoolRole && user.Role != models.AdminRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	var update models.SchoolUpdate
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.School.Update(r.Context(), user, update); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(statusResponse{Status: "OK"}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary getAllSchools
// @Security ApiKeyAuth
// @Tags school
// @Produce json
// @Success 200 {object} models.FullSchool
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/school/all [get]
func (h *Handler) getAllSchools(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.AdminRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	schools, err := h.service.School.GetAllForAdmin(r.Context())
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(schools); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary deleteSchool
// @Security ApiKeyAuth
// @Tags school
// @Produce json
// @Param id path int true "school id"
// @Success 200 {object} object
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/school/delete/{id} [delete]
func (h *Handler) deleteSchool(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.AdminRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	schoolId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		h.errorHandler(w, r, http.StatusNotFound, err.Error())
		return
	}

	err = h.service.School.Delete(r.Context(), uint(schoolId))
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(statusResponse{Status: "OK"}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}
