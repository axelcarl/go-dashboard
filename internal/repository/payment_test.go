package repository

import (
	"go-dashboard/internal/database"
	"testing"
)

func TestNewPaymentRepository(t *testing.T) {
	service := database.New()

	repository := NewPaymentRepository(service.DB())
	if repository == nil {
		t.Fatal("NewPaymentRepository() returned nil")
	}
}
