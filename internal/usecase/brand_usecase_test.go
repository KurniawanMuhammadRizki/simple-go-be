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

func TestCreateBrand(t *testing.T) {
	// Use temporary file instead of in-memory
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect database: %v", err)
	}

	// Explicitly create table
	err = db.Exec(`CREATE TABLE brands (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    name TEXT,
	    created_at DATETIME,
	    updated_at DATETIME
	)`).Error
	if err != nil {
		t.Fatalf("Failed to create table: %v", err)
	}

	log := logrus.New()
	repo := repository.NewBrandRepository(log)
	uc := NewBrandUsecase(repo, log, db)

	req := &model.CreateBrandRequest{Name: "Test Brand"}
	resp, err := uc.CreateBrand(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Test Brand", resp.Name)
}
