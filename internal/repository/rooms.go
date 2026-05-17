package repository

import (
	"booking-system/lvl1/internal/config"
	"booking-system/lvl1/internal/model"
)

type RoomRepository struct {
	db *config.DB
}

func NewRoomRepository(db *config.DB) *RoomRepository {
	return &RoomRepository{db: db}
}

func (r *RoomRepository) CreateRoom(room *model.Room) error {
	return r.db.Create(room).Error
}

func (r *RoomRepository) GetRoomsByHotelID(hotelID string) ([]model.Room, error) {
	var rooms []model.Room
	result := r.db.Where("hotel_id = ?", hotelID).Find(&rooms)
	return rooms, result.Error
}

func (r *RoomRepository) GetRoomByID(id string) (*model.Room, error) {
	var room model.Room
	result := r.db.Where("id = ?", id).First(&room)
	if result.Error != nil {
		return nil, result.Error
	}
	return &room, nil
}

func (r *RoomRepository) UpdateRoom(room *model.Room) error {
	return r.db.Save(room).Error
}

func (r *RoomRepository) DeleteRoom(id string) error {
	return r.db.Delete(&model.Room{}, "id = ?", id).Error
}
