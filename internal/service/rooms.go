package service

import (
	"booking-system/lvl1/internal/model"
	"booking-system/lvl1/internal/repository"
)

type RoomService struct {
	repo *repository.RoomRepository
}

func NewRoomService(repo *repository.RoomRepository) *RoomService {
	return &RoomService{repo: repo}
}

func (s *RoomService) CreateRoom(hotelID string, req *model.CreateRoomRequest) error {
	room := &model.Room{
		HotelID:      hotelID,
		RoomNumber:   req.RoomNumber,
		Type:         req.Type,
		Description:  req.Description,
		Price:        req.Price,
		Capacity:     req.Capacity,
		BedCount:     req.BedCount,
		BedType:      req.BedType,
		HasWifi:      req.HasWifi,
		HasAircon:    req.HasAircon,
		HasTV:        req.HasTV,
		HasBreakfast: req.HasBreakfast,
	}
	return s.repo.CreateRoom(room)
}

func (s *RoomService) GetRoomsByHotel(hotelID string) ([]model.Room, error) {
	return s.repo.GetRoomsByHotelID(hotelID)
}

func (s *RoomService) GetRoomByID(id string) (*model.Room, error) {
	return s.repo.GetRoomByID(id)
}

func (s *RoomService) UpdateRoom(id string, req *model.UpdateRoomRequest) error {
	room, err := s.repo.GetRoomByID(id)
	if err != nil {
		return err
	}
	if req.RoomNumber != "" {
		room.RoomNumber = req.RoomNumber
	}
	if req.Type != "" {
		room.Type = req.Type
	}
	if req.Description != "" {
		room.Description = req.Description
	}
	if req.Capacity != 0 {
		room.Capacity = req.Capacity
	}
	if req.BedCount != 0 {
		room.BedCount = req.BedCount
	}
	if req.BedType != "" {
		room.BedType = req.BedType
	}
	if req.Price != 0 {
		room.Price = req.Price
	}
	room.HasWifi = req.HasWifi
	room.HasAircon = req.HasAircon
	room.HasTV = req.HasTV
	room.HasBreakfast = req.HasBreakfast
	return s.repo.UpdateRoom(room)
}

func (s *RoomService) DeleteRoom(id string) error {
	return s.repo.DeleteRoom(id)
}
