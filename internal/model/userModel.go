package model

import (
	"time"
	"uuid"
)

type UserModel struct {
	ID uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Username string `gorm:"unique"`
	Password string 
	CreatedAt time.Time 
	UpdatedAt time.Time 
	Todos []TodoModel `gorm:"foreignKey:UserID"`
}