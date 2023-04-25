package http

import (
	"book-school/internal/models"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

// @Summary getAllBooksForTest
// @Security ApiKeyAuth
// @Tags book
// @Accept json
// @Produce json
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/books [get]
func (h *Handler) getAllBooksForTest(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.AdminRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	books, err := h.service.Book.GetAllForTest(r.Context())
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(books); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary createBook
// @Security ApiKeyAuth
// @Tags book
// @Accept json
// @Produce json
// @Param newBook body models.Book true "book info"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/books/create [post]
func (h *Handler) createBook(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.AdminRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	var input models.Book
	jsonDoc := r.FormValue("document")
	file, _, err := r.FormFile("preview")
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if err = json.Unmarshal([]byte(jsonDoc), &input); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.service.Book.CreateBook(&input, file); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}
	if err := json.NewEncoder(w).Encode(statusResponse{Status: "OK"}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary deleteBook
// @Security ApiKeyAuth
// @Tags book
// @Accept json
// @Produce json
// @Param book_id path int true "Book id"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/books/delete/{book_id} [delete]
func (h *Handler) deleteBook(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	if user.Role != models.AdminRole {
		h.errorHandler(w, r, http.StatusUnauthorized, "invalid role")
		return
	}
	bookId, err := strconv.Atoi(mux.Vars(r)["book_id"])
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
	}
	if bookId < 1 {
		h.errorHandler(w, r, http.StatusBadRequest, "not correct id")
		return
	}
	err = h.service.Book.DeleteBook(bookId)
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
	}
	if err = json.NewEncoder(w).Encode(statusResponse{Status: "OK"}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary getAllBooks
// @Security ApiKeyAuth
// @Tags book
// @Produce json
// @Success 200 {object} []models.BooksStruct
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/books/all/ [get]
func (h *Handler) getAllBooks(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)
	log.Println("lol")
	books, err := h.service.Book.GetAll(r.Context(), user)
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}
	log.Println(books)
	if err = json.NewEncoder(w).Encode(books); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary getBookById
// @Security ApiKeyAuth
// @Tags book
// @Accept json
// @Produce json
// @Param id path int true "Book id"
// @Success 200 {object} models.Book
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/books/{id} [get]
func (h *Handler) getBookById(w http.ResponseWriter, r *http.Request) {
	bookId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}
	if bookId < 1 {
		h.errorHandler(w, r, http.StatusBadRequest, "not correct id")
		return
	}
	book, err := h.service.Book.GetJsonBook(bookId)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			h.errorHandler(w, r, http.StatusBadRequest, err.Error())
			return
		}
	}
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	if err = encoder.Encode(book); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary totalPages
// @Security ApiKeyAuth
// @Tags book
// @Accept json
// @Produce json
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/books/total [get]
func (h *Handler) totalPages(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)

	totalPages, err := h.service.Book.TotalPages(r.Context(), user)
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(statusResponse{Status: strconv.Itoa(totalPages)}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary completeBook
// @Security ApiKeyAuth
// @Tags book
// @Produce json
// @Param id path int true "Book id"
// @Success 200 {object} object
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/books/complete/{id} [post]
func (h *Handler) completeBook(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)
	if user.Role != models.StudentRole {
		h.errorHandler(w, r, http.StatusBadRequest, "invalid role")
		return
	}

	bookId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}
	if bookId < 1 {
		h.errorHandler(w, r, http.StatusBadRequest, "not correct id")
		return
	}

	points, err := h.service.Book.Complete(r.Context(), uint(bookId), user.Id)
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(points); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary completeBook
// @Security ApiKeyAuth
// @Tags book
// @Produce json
// @Param book_id path int true "book id"
// @Success 200 {object} object
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/audio/complete/{book_id} [post]
func (h *Handler) completeAudio(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(models.UserCtx).(models.User)
	if user.Role != models.StudentRole {
		h.errorHandler(w, r, http.StatusBadRequest, "invalid role")
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

	points, err := h.service.Audio.Complete(r.Context(), uint(bookId), user.Id)
	if err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(points); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}
