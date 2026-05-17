package repository

import (
	"booking-system/lvl1/internal/config"
	"booking-system/lvl1/internal/model"
)

type HotelRepository struct {
	db *config.DB
}

func NewHotelRepository(db *config.DB) *HotelRepository {
	return &HotelRepository{db: db}
}

func (r *HotelRepository) CreateHotel(hotel *model.Hotel) error {
	return r.db.Create(hotel).Error
}

func (r *HotelRepository) GetHotelByID(id string) (*model.Hotel, error) {
	var hotel model.Hotel
	if err := r.db.First(&hotel, id).Omit("created_at", "updated_at").Error; err != nil {
		return nil, err
	}
	return &hotel, nil
}

func (r *HotelRepository) GetAllHotels() ([]model.Hotel, error) {
	var hotels []model.Hotel
	if err := r.db.Find(&hotels).Omit("created_at", "updated_at").Error; err != nil {
		return nil, err
	}
	return hotels, nil
}

func (r *HotelRepository) UpdateHotel(hotel *model.Hotel) error {
	return r.db.Save(hotel).Error
}
