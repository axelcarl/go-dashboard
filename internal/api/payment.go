package api

import (
	"go-dashboard/internal/api/mapper"
	"go-dashboard/internal/api/response"
	interfaces "go-dashboard/internal/application/interface"
	"go-dashboard/internal/application/query"
	"go-dashboard/internal/application/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type PaymentHandler struct {
	service interfaces.PaymentService
}

func NewPaymentHandler(r *chi.Mux, service *service.PaymentService) *PaymentHandler {
	handler := PaymentHandler{service: service}

	r.Route("/payment", func(r chi.Router) {
		r.Get("/{id}", handler.GetPaymentByID)
	})

	return &handler
}

func (ph *PaymentHandler) GetPaymentByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := response.ErrorResponse{Error: "Invalid identifier"}
		render.JSON(w, r, response)
		return
	}

	payment, err := ph.service.FindPaymentByID(&query.GetPaymentByIDQuery{ID: id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := response.ErrorResponse{Error: "Failed to fetch payment"}
		render.JSON(w, r, response)
		return
	}

	if payment == nil {
		w.WriteHeader(http.StatusNotFound)
		response := response.ErrorResponse{Error: "Payment not found"}
		render.JSON(w, r, response)
		return
	}

	response := mapper.ToPaymentResponse(payment.Result)
	render.JSON(w, r, response)
}
