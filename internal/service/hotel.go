package service

import (
	"booking-system/lvl1/internal/model"
	"booking-system/lvl1/internal/repository"
)

type HotelService struct {
	repo *repository.HotelRepository
}

func NewHotelService(repo *repository.HotelRepository) *HotelService {
	return &HotelService{repo: repo}
}

func (s *HotelService) CreateHotel(req *model.CreateHotelRequest) error {
	hotel := &model.Hotel{
		Name:        req.Name,
		Description: req.Description,
		Address:     req.Address,
		City:        req.City,
		Rating:      req.Rating,
		UserID:      req.UserID,
	}
	return s.repo.CreateHotel(hotel)
}

func (s *HotelService) GetHotelByID(id string) (*model.Hotel, error) {
	return s.repo.GetHotelByID(id)
}

func (s *HotelService) GetAllHotels() ([]model.Hotel, error) {
	return s.repo.GetAllHotels()
}

func (s *HotelService) UpdateHotel(id string, req *model.UpdateHotelRequest) error {
	hotel, err := s.repo.GetHotelByID(id)
	if err != nil {
		return err
	}
	hotel.Name = req.Name
	hotel.Description = req.Description
	hotel.Address = req.Address
	hotel.City = req.City
	hotel.Rating = req.Rating
	return s.repo.UpdateHotel(hotel)
}
