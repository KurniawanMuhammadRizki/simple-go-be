package usecase

import (
	"context"
	"testing"

	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/model"
	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/repository"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateCustomer(t *testing.T) {
	// Open in-memory database for testing
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	// Explicitly create customers table in database
	err = db.Exec(`CREATE TABLE customers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		point_balance BIGINT,
		created_at DATETIME,
		updated_at DATETIME
	)`).Error
	if err != nil {
		t.Fatalf("Failed to create table: %v", err)
	}

	// Initialize logger, repository, and usecase
	log := logrus.New()
	repo := repository.NewCustomerRepository(log)
	uc := NewCustomerUsecase(repo, log, db)

	// Prepare test request data
	req := &model.CreateCustomerRequest{
		Name: "Test Customer",
	}

	// Call the CreateCustomer method
	resp, err := uc.CreateCustomer(context.Background(), req)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Test Customer", resp.Name)

}
