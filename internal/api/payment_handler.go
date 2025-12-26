package api

import (
	"go-dashboard/internal/api/mapper"
	"go-dashboard/internal/api/response"
	"go-dashboard/internal/application/query"
	"go-dashboard/internal/application/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type PaymentHandler struct {
	service *service.PaymentService
}

func NewPaymentHandler(r *chi.Mux, service *service.PaymentService) *PaymentHandler {
	handler := PaymentHandler{service: service}

	r.Route("/payment", func(r chi.Router) {
		r.Get("/{id}", handler.GetProductById)
	})

	return &handler
}

func (ph *PaymentHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(400)
		response := response.ErrorResponse{Error: "Invalid identifier"}
		render.JSON(w, r, response)
		return
	}

	payment, err := ph.service.FindPaymentByID(&query.GetPaymentByIDQuery{ID: id})
	if err != nil {
		w.WriteHeader(500)
		response := response.ErrorResponse{Error: "Failed to fetch payment"}
		render.JSON(w, r, response)
		return
	}

	if payment == nil {
		w.WriteHeader(404)
		response := response.ErrorResponse{Error: "Payment not found"}
		render.JSON(w, r, response)
		return
	}

	response := mapper.ToPaymentResponse(payment.Result)
	render.JSON(w, r, response)
}
