package model

type Category struct {
	Index    string `json:"index" gorm:"type:char(36); not null"`
	Label    string `json:"label" gorm:"type:varchar(50);not null"`
	MainMenu string `json:"mainmenu" gorm:"type:varchar(50);not null"`
}
