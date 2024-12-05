package repository

import (
	"errors"

	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BrandRepository struct {
	Log *logrus.Logger
}

func NewBrandRepository(log *logrus.Logger) *BrandRepository {
	return &BrandRepository{
		Log: log,
	}
}

func (b *BrandRepository) Save(db *gorm.DB, brand *entity.Brand) (*entity.Brand, error) {
	err := db.Create(brand).Error
	if err != nil {
		b.Log.Error(err)
		return nil, err
	}

	return brand, nil
}

// GetByID implements BrandRepository.
func (b *BrandRepository) GetByID(db *gorm.DB, id int64) (*entity.Brand, error) {
	var brand entity.Brand
	err := db.Where("id = ?", id).First(&brand).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			b.Log.WithField("id", id).Error("brand not found")
			return nil, err
		}
		b.Log.Error(err)
		return nil, err
	}
	return &brand, nil
}

// Update implements BrandRepository.
func (b *BrandRepository) Update(db *gorm.DB, brand *entity.Brand) (*entity.Brand, error) {
	err := db.Where("id = ?", brand.ID).Updates(brand).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			b.Log.WithField("id", brand.ID).Error("brand not found")
			return nil, err
		}
		b.Log.Error(err)
		return nil, err
	}

	// Fetch the updated Brand
	return b.GetByID(db, brand.ID)
}

// Delete implements BrandRepository.
func (b *BrandRepository) Delete(db *gorm.DB, id int64) error {
	result := db.Model(&entity.Brand{}).
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
