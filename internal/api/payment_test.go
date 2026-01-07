package api

import (
	"go-dashboard/internal/application/common"
	"go-dashboard/internal/application/mutation"
	"go-dashboard/internal/application/query"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
)

type mockPaymentService struct {
	findFn         func(q *query.GetPaymentByIDQuery) (*query.GetPaymentByIdQueryResult, error)
	findMultipleFn func() (*query.GetPaymentsQueryResult, error)
	create         func(m *mutation.CreatePaymentMutation) (*mutation.CreatePaymentMutationResult, error)
}

func (m *mockPaymentService) FindPaymentByID(q *query.GetPaymentByIDQuery) (*query.GetPaymentByIdQueryResult, error) {
	return m.findFn(q)
}

func (m *mockPaymentService) List() (*query.GetPaymentsQueryResult, error) {
	return m.findMultipleFn()
}

func (m *mockPaymentService) Create(mut *mutation.CreatePaymentMutation) (*mutation.CreatePaymentMutationResult, error) {
	return m.Create(mut)
}

func TestGetPaymentByID_Ok(t *testing.T) {
	mockService := &mockPaymentService{
		findFn: func(q *query.GetPaymentByIDQuery) (*query.GetPaymentByIdQueryResult, error) {
			if q.ID != 2 {
				t.Fatalf("expected ID 2, got %d", q.ID)
			}

			return &query.GetPaymentByIdQueryResult{
				Result: &common.PaymentResult{
					ID:        2,
					Sender:    "John",
					Recipient: "Steven",
					Amount:    100,
					CreatedAt: time.Date(2025, 12, 25, 10, 28, 53, 104233000, time.UTC),
					UpdatedAt: time.Date(2025, 12, 25, 10, 28, 53, 104233000, time.UTC),
				},
			}, nil
		},
	}

	ph := &PaymentHandler{
		service: mockService,
	}

	r := chi.NewRouter()
	r.Get("/payments/{id}", ph.GetPaymentByID)

	req := httptest.NewRequest(http.MethodGet, "/payments/2", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	expected := `{"id":2,"sender":"John","recipient":"Steven","amount":100,"updatedAt":"2025-12-25T10:28:53.104233Z","createdAt":"2025-12-25T10:28:53.104233Z"}`
	if strings.TrimSpace(rec.Body.String()) != expected {
		t.Fatalf("didn't get expected response, got: %s", rec.Body.String())
	}
}

func TestGetPaymentByID_InvalidID(t *testing.T) {
	ph := &PaymentHandler{}

	r := chi.NewRouter()
	r.Get("/payments/{id}", ph.GetPaymentByID)

	req := httptest.NewRequest(http.MethodGet, "/payments/abc", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}

	expected := `{"error":"Invalid identifier"}`
	if strings.TrimSpace(rec.Body.String()) != expected {
		t.Fatalf("didn't get expected response, got %s", rec.Body.String())
	}
}
