package api

import (
	"go-dashboard/internal/api/mapper"
	"go-dashboard/internal/api/request"
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

	r.Route("/payments", func(r chi.Router) {
		r.Get("/", handler.GetPayments)
		r.Post("/", handler.CreatePayment)
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

func (ph *PaymentHandler) GetPayments(w http.ResponseWriter, r *http.Request) {
	payments, err := ph.service.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := response.ErrorResponse{Error: "Something went wrong when fetching payments"}
		render.JSON(w, r, response)
		return
	}

	response := mapper.ToPaymentListResponse(payments.Result)
	render.JSON(w, r, response)
}

func (ph *PaymentHandler) CreatePayment(w http.ResponseWriter, r *http.Request) {
	var body request.PaymentRequest
	err := render.DecodeJSON(r.Body, &body)
	if err != nil {
		WriteErrorResponse(w, r, http.StatusBadRequest, "Request body doesn't follow schema")
		return
	}

	mut, err := body.ToCreatePaymentMutation()
	if err != nil {
		switch err {
		case request.ErrSenderEmpty:
			WriteErrorResponse(w, r, http.StatusBadRequest, "sender can't be empty")

		case request.ErrRecipientEmpty:
			WriteErrorResponse(w, r, http.StatusBadRequest, "recipient can't be empty")

		case request.ErrInvalidAmount:
			WriteErrorResponse(w, r, http.StatusBadRequest, "amount must be greater than 0")
		}

		return
	}

	payment, err := ph.service.Create(mut)
	if err != nil {
		WriteErrorResponse(w, r, http.StatusInternalServerError, "Failed to create payment")
		return
	}

	response := mapper.ToPaymentResponse(payment.Result)
	render.JSON(w, r, response)
}
