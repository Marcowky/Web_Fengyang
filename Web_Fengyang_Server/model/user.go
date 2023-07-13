package model

// TODO 添加用户级别，用以区分用户的权限以及相应功能
type User struct {
	ID          uint   `gorm:"primaryKey;not null"`
	UserName    string `gorm:"varchar(20);not null"`
	PhoneNumber string `gorm:"varchar(20);not null;unique"`
	CreatedAt   Time   `gorm:"type:timestamp"`
	Password    string `gorm:"size:255;not null"`
	UserType    string `gorm:"varchar(20);not null"`
	Status      bool   `gorm:"not null"`
}
