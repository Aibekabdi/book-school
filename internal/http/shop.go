package http

import (
	"book-school/internal/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// @Summary buyBody
// @Security ApiKeyAuth
// @Tags shop
// @Produce json
// @Param body_id path int true "body id"
// @Success 200 {object} object
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/shop/buy/{body_id} [post]
func (h *Handler) buyBody(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.StudentRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	bodyId, err := strconv.Atoi(mux.Vars(r)["body_id"])
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	points, err := h.service.Shop.Buy(r.Context(), user.Id, uint(bodyId))
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(points); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary getAllBody
// @Security ApiKeyAuth
// @Tags shop
// @Produce json
// @Success 200 {object} models.ShopInfo
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/shop/all [get]
func (h *Handler) getAllBody(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.StudentRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	info, err := h.service.Shop.GetAll(r.Context(), user.Id)
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(info); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary createNewBody
// @Security ApiKeyAuth
// @Tags shop
// @Accept mpfd
// @Produce json
// @Param image formData file true "body image"
// @Param imageIcon formData file true "body image icon"
// @Param part formData string true "body part"
// @Param name formData string true "body name"
// @Param price formData int true "body price"
// @Success 200 {object} models.Body
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/shop/create [post]
func (h *Handler) createNewBody(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.AdminRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	image, ok := r.MultipartForm.File["image"]
	if !ok {
		h.errorHandler(w, r, http.StatusBadRequest, "image field not found")
		return
	}

	imageIcon, ok := r.MultipartForm.File["imageIcon"]
	if !ok {
		h.errorHandler(w, r, http.StatusBadRequest, "image icon field not found")
		return
	}

	part, ok := r.Form["part"]
	if !ok {
		h.errorHandler(w, r, http.StatusBadRequest, "part field not found")
		return
	}

	name, ok := r.Form["name"]
	if !ok {
		h.errorHandler(w, r, http.StatusBadRequest, "name field not found")
		return
	}

	price, ok := r.Form["price"]
	if !ok {
		h.errorHandler(w, r, http.StatusBadRequest, "price field not found")
		return
	}

	body, err := h.service.Shop.Create(r.Context(), image[0], imageIcon[0], price[0], name[0], part[0])
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(body); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}
