package models

type Email struct {
	Id         int `json:"id,omitempty" gorm:"primaryKey"`
	User       User
	UserID     string `json:"user_id"`
	IsPrimary  bool   `json:"is_primary"`
	IsVerified bool   `json:"is_verified"`
	Email      string `json:"email"`
}

type EmailVerification struct {
	Id               int    `json:"id,omitempty" gorm:"primaryKey"`
	Email            Email  `gorm:"constraint:onUpdate:CASCADE,onDelete:CASCADE;"`
	EmailID          int    `json:"email_id"`
	VerificationCode string `json:"verification_code"`
	ExpiresAt        int64  `json:"expires_at"`
}

type Mask struct {
	Mask              string `json:"mask" gorm:"primaryKey"`
	Enabled           bool   `json:"enabled"`
	Email             Email  `gorm:"foreignKey:ForwardTo"`
	ForwardTo         int    `json:"forward_to"`
	User              User
	UserID            string `json:"user_id"`
	MessagesReceived  int    `json:"messages_received" gorm:"default:0"`
	MessagesForwarded int    `json:"messages_forwarded" gorm:"default:0"`
}
