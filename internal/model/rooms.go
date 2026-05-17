package model

import "time"

type RoomType string

const (
	RoomTypeSingle RoomType = "Deluxe"
	RoomTypeDouble RoomType = "Standard"
	RoomTypeSuite  RoomType = "Family"
)

type Room struct {
	ID           string   `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	HotelID      string   `gorm:"not null;index"`
	Hotel        Hotel    `gorm:"foreignKey:HotelID"`
	RoomNumber   string   `gorm:"not null;size:50"`
	Type         RoomType `gorm:"not null"`
	Description  string   `gorm:"type:text"`
	Capacity     int8     `gorm:"not null;default:1"`
	BedCount     int8     `gorm:"not null;default:1"`
	BedType      string   `gorm:"size:50"`
	HasWifi      bool     `gorm:"default:false"`
	HasAircon    bool     `gorm:"default:false"`
	HasTV        bool     `gorm:"default:false"`
	HasBreakfast bool     `gorm:"default:false"`
	Price        float64  `gorm:"type:decimal(10,2);not null"`
	Available    bool     `gorm:"not null;default:true"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type CreateRoomRequest struct {
	RoomNumber   string   `json:"room_number" binding:"required,min=1,max=50"`
	Type         RoomType `json:"type" binding:"required,oneof=Deluxe Standard Family"`
	Description  string   `json:"description,omitempty"`
	Capacity     int8     `json:"capacity" binding:"required,gt=0"`
	BedCount     int8     `json:"bed_count" binding:"required,gt=0"`
	BedType      string   `json:"bed_type,omitempty" binding:"max=50"`
	HasWifi      bool     `json:"has_wifi"`
	HasAircon    bool     `json:"has_aircon"`
	HasTV        bool     `json:"has_tv"`
	HasBreakfast bool     `json:"has_breakfast"`
	Price        float64  `json:"price" binding:"required,gt=0"`
}

type UpdateRoomRequest struct {
	RoomNumber   string   `json:"room_number" binding:"omitempty,min=1,max=50"`
	Type         RoomType `json:"type" binding:"omitempty,oneof=Deluxe Standard Family"`
	Description  string   `json:"description,omitempty"`
	Capacity     int8     `json:"capacity" binding:"omitempty,gt=0"`
	BedCount     int8     `json:"bed_count" binding:"omitempty,gt=0"`
	BedType      string   `json:"bed_type,omitempty" binding:"max=50"`
	HasWifi      bool     `json:"has_wifi"`
	HasAircon    bool     `json:"has_aircon"`
	HasTV        bool     `json:"has_tv"`
	HasBreakfast bool     `json:"has_breakfast"`
	Price        float64  `json:"price" binding:"omitempty,gt=0"`
}
