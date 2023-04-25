package http

import (
	"book-school/internal/models"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
)

// @Summary signIn
// @Tags auth
// @Accept json
// @Produce json
// @Param signIn body models.SignInInput true "sign in info"
// @Success 200 {object} object
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var (
		signIn models.SignInInput
		err    error
		token  string
	)
	if err = json.NewDecoder(r.Body).Decode(&signIn); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}
	token, err = h.service.Auth.SignIn(r.Context(), signIn.Name, signIn.Password, signIn.Who)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		default:
			h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		}
		return
	}
	if err = json.NewEncoder(w).Encode(map[string]interface{}{"token": token}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}
