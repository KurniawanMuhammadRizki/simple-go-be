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

func TestCreateVoucher(t *testing.T) {
	// Use temporary file instead of in-memory
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect database: %v", err)
	}

	err = db.Exec(`CREATE TABLE vouchers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		brand_id INTEGER,
		name TEXT,
		cost_in_points INTEGER,
		created_at DATETIME,
		updated_at DATETIME,
		FOREIGN KEY (brand_id) REFERENCES brands(id)
	)`).Error
	if err != nil {
		t.Fatalf("Failed to create vouchers table: %v", err)
	}

	// Insert sample brand data
	db.Exec(`INSERT INTO brands (name, created_at, updated_at) VALUES ('Brand AB', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`)

	// Initialize repository and usecase
	log := logrus.New()
	repo := repository.NewVoucherRepository(log)
	uc := NewVoucherUsecase(repo, log, db)

	// Prepare request to create a voucher
	req := &model.CreateVoucherRequest{
		Name:        "Test Voucher",
		BrandID:     1,
		CostInPoint: 100,
	}

	// Call the usecase to create the voucher
	resp, err := uc.CreateVoucher(context.Background(), req)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Test Voucher", resp.Name)
	assert.Equal(t, int64(1), resp.BrandID)
	assert.Equal(t, int64(100), resp.CostInPoint)
}
