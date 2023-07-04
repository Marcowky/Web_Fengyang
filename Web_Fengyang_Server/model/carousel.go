package model

// import "github.com/jinzhu/gorm"

/* model/carousel.go */
type Carousel struct {
	ID       uint   `gorm:"primaryKey;not null"`
	Order    string `gorm:"varchar(20)"`
	Category string `gorm:"varchar(20);not null"`
	Url      string `gorm:"type:text;not null"`
}
