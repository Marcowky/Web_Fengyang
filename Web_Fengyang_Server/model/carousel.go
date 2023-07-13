package model

type Carousel struct {
	ID       uint   `gorm:"primaryKey;not null"`
	Iorder   string `gorm:"varchar(20)"`
	Category string `gorm:"varchar(20);not null"`
	Url      string `gorm:"type:text;not null"`
}
