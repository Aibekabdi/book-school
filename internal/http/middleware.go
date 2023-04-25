package http

import (
	"book-school/internal/models"
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func (h *Handler) corsMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) loggingMiddleWare(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Printf("%s %s [%s]\t%s%s - 200 - OK\n", time.Now().Format("2006/01/02 15:04:05"), r.Proto, r.Method, r.Host, r.RequestURI)
		f.ServeHTTP(w, r)
	})
}

func (h *Handler) userIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header, ok := r.Header["Authorization"]
		if !ok {
			h.errorHandler(w, r, http.StatusUnauthorized, "empty auth header")
			return
		}
		headerParts := strings.Split(header[0], " ")
		if len(headerParts) != 2 {
			h.errorHandler(w, r, http.StatusUnauthorized, "invalid auth header")
			return
		}
		user, err := h.service.Auth.ParseToken(headerParts[1])
		if err != nil {
			h.errorHandler(w, r, http.StatusUnauthorized, err.Error())
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), models.UserCtx, user)))
	})
}
