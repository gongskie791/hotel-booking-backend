package model

import "time"

type Hotel struct {
	ID          string       `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name        string       `gorm:"not null;size:255" binding:"required,min=2,max=255"`
	Description string       `gorm:"size:255"`
	Address     string       `gorm:"not null;size:255" binding:"required,min=2,max=255"`
	City        string       `gorm:"not null;size:255" binding:"required,min=2,max=255"`
	Rating      float32      `gorm:"type:numeric(10,2)"`
	Images      []HotelImage `gorm:"foreignKey:HotelID"`
	UserID      string       `gorm:"type:uuid;not null"` // links to User
	User        User         `gorm:"foreignKey:UserID"`  // belongs to User
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type HotelImage struct {
	ID      string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	HotelID string `gorm:"not null;index"`
	URL     string `gorm:"not null;size:255" binding:"required,url"`
}

type CreateHotelRequest struct {
	Name        string   `json:"name" binding:"required,min=2,max=255"`
	Description string   `json:"description" binding:"max=255"`
	Address     string   `json:"address" binding:"required,min=2,max=255"`
	City        string   `json:"city" binding:"required,min=2,max=255"`
	Rating      float32  `json:"rating" binding:"gte=0,lte=5"`
	UserID      string   `json:"user_id" binding:"required,uuid"`
	Images      []string `json:"images,omitempty" binding:"dive,required,url"`
}

type UpdateHotelRequest struct {
	Name        string   `json:"name" binding:"required,min=2,max=255"`
	Description string   `json:"description" binding:"max=255"`
	Address     string   `json:"address" binding:"required,min=2,max=255"`
	City        string   `json:"city" binding:"required,min=2,max=255"`
	Rating      float32  `json:"rating" binding:"gte=0,lte=5"`
	Images      []string `json:"images,omitempty" binding:"dive,required,url"`
}
