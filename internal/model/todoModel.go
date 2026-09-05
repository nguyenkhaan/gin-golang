package model

import (
	"time"
	"uuid"
	"gorm.io/gorm"
)

type TodoModel struct {
	gorm.Model 
	ID uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Title string 
	Content string `gorm:"default:''"`
	CreatedAt time.Time 
	UpdatedAt time.Time 
	DeletedAt gorm.DeletedAt `gorm:"index"`
	//Define the foreign key 
	UserID uuid.UUID //Foreign Key ID field
	User UserModel `gorm:"foreignKey:UserID"`
}