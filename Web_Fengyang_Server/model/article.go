package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Article struct {
	ID          uuid.UUID `json:"id" gorm:"type:char(36);primary_key;"`
	UserId      uint      `json:"user_id" gorm:"not null"`
	CategoryId  uint      `json:"category_id" gorm:"not null"`
	Title       string    `json:"title" gorm:"type:varchar(50);not null"`
	Content     string    `json:"content" gorm:"type:text;not null"`
	HeadImage   string    `json:"head_image"`
	ArticleType string    `json:"article_type" gorm:"varchar(20);not null"`
	CreatedAt   Time      `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt   Time      `json:"updated_at" gorm:"type:timestamp"`
}

type ArticleInfo struct {
	ID         string `json:"id"`
	CategoryId uint   `json:"category_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	HeadImage  string `json:"head_image"`
	CreatedAt  Time   `json:"created_at"`
}

// BeforeCreate 在创建文章之前将id赋值
func (a *Article) BeforeCreate(s *gorm.Scope) error {
	return s.SetColumn("ID", uuid.NewV4())
}

// 用于表单验证
type CreateArticleRequest struct {
	// 加上binding用于表单验证
	CategoryId  uint   `json:"category_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	HeadImage   string `json:"head_image"`
	ArticleType string `json:"article_type" gorm:"varchar(20);not null"`
}
