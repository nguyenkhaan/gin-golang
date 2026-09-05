package model

import (
	"time"
	"uuid"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	ID uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Username string `gorm:"unique"`
	Password string 
	CreatedAt time.Time 
	UpdatedAt time.Time 
	Todos []TodoModel `gorm:"foreignKey:UserID"`
}