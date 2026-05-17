package model

type UserType string

const (
	UserTypeUser  UserType = "user"
	UserTypeAdmin UserType = "admin"
)

type User struct {
	ID              string           `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Username        string           `gorm:"unique;not null" binding:"required,min=2,max=100"`
	Password        string           `gorm:"not null"`
	Role            UserType         `gorm:"not null;default:'user'"`
	PersonalDetails *PersonalDetails `gorm:"foreignKey:UserID"` // has one
}

type PersonalDetails struct {
	ID     string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID string `gorm:"type:uuid;not null;uniqueIndex"` // links to User
	Name   string `gorm:"not null;size:100" binding:"required,min=2,max=100"`
	Age    int8   `gorm:"not null;default:18" binding:"required,min=18,max=120"`
	Gender string `gorm:"size:20" binding:"omitempty,oneof=male female other"`
	Email  string `gorm:"unique;not null"`
	Phone  string `gorm:"not null;min=10;max=15"`
}

type CreateUserRequest struct {
	Username string   `json:"username" binding:"required,min=4,max=100"`
	Password string   `json:"password" binding:"required,min=8,max=100"`
	Role     UserType `json:"role,omitempty" binding:"omitempty,oneof=user admin"`
}
