package models

import "time"

type Product struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name" form:"name"`
	Description string    `json:"description" form:"description"`
	CreatedAt   time.Time `json:"created_at" gorm:"type:TIMESTAMP(6)"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"type:TIMESTAMP(6)"`
}

func (Product) TableName() string {
	return "product"
}
