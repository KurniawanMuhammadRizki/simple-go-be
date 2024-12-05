package repository

import (
	"errors"

	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type VoucherRepository struct {
	Log *logrus.Logger
}

func NewVoucherRepository(log *logrus.Logger) *VoucherRepository {
	return &VoucherRepository{
		Log: log,
	}
}

func (b *VoucherRepository) Save(db *gorm.DB, voucher *entity.Voucher) (*entity.Voucher, error) {
	err := db.Create(voucher).Error
	if err != nil {
		b.Log.Error(err)
		return nil, err
	}

	return voucher, nil
}

// GetByID implements VoucherRepository.
func (b *VoucherRepository) GetByID(db *gorm.DB, id int64) (*entity.Voucher, error) {
	var voucher entity.Voucher
	err := db.Where("id = ?", id).First(&voucher).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			b.Log.WithField("id", id).Error("voucher not found")
			return nil, err
		}
		b.Log.Error(err)
		return nil, err
	}
	return &voucher, nil
}

func (b *VoucherRepository) GetAllByBrand(db *gorm.DB, brandID int64) ([]entity.Voucher, error) {
	var vouchers []entity.Voucher

	err := db.Where("brand_id = ?", brandID).Find(&vouchers).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			b.Log.WithField("brand_id", brandID).Error("no vouchers found for the given brand")
			return nil, err
		}
		b.Log.Error(err)
		return nil, err
	}

	return vouchers, nil
}

// Update implements VoucherRepository.
func (b *VoucherRepository) Update(db *gorm.DB, voucher *entity.Voucher) (*entity.Voucher, error) {
	err := db.Where("id = ?", voucher.ID).Updates(voucher).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			b.Log.WithField("id", voucher.ID).Error("Voucher not found")
			return nil, err
		}
		b.Log.Error(err)
		return nil, err
	}

	// Fetch the updated Voucher
	return b.GetByID(db, voucher.ID)
}

// Delete implements VoucherRepository.
func (b *VoucherRepository) Delete(db *gorm.DB, id int64) error {
	result := db.Model(&entity.Voucher{}).
		Where("id = ? ", id)

	if result.Error != nil {
		b.Log.Error(result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
