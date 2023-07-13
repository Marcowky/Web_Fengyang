package model

// TODO 添加用户级别，用以区分用户的权限以及相应功能
type Category struct {
	Order    uint   `gorm:"not null"`
	Label    string `gorm:"varchar(20);not null"`
	PageName string `gorm:"varchar(20);not null"`
}
