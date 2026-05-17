package model

// type UserToken struct {
// 	id           string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
// 	UserID       string `gorm:"type:uuid;not null;uniqueIndex"`
// 	AccessToken  string `gorm:"not null"`
// 	RefreshToken string `gorm:"not null"`
// 	User         User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
// }

type LoginUserInfo struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type LoginResponse struct {
	Token string        `json:"token"`
	User  LoginUserInfo `json:"user"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
