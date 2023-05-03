package model

type Category struct {
	ID           *int   `json:"id" gorm:"type:char(36);primary_key;"`
	CategoryName string `json:"name" gorm:"type:varchar(50);not null"`
}
