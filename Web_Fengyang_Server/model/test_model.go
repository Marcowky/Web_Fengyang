package model

import "github.com/jinzhu/gorm"

type TestModel struct {
	gorm.Model
	TestData int
}