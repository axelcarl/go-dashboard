package api

import (
	"go-dashboard/internal/api/response"
	"net/http"

	"github.com/go-chi/render"
)

func WriteErrorResponse(w http.ResponseWriter, r *http.Request, status int, message string) {
	w.WriteHeader(http.StatusBadRequest)
	response := response.ErrorResponse{Error: message}
	render.JSON(w, r, response)
}
