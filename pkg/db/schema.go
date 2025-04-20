package db

import (
	"time"

	"gorm.io/gorm"
)

// 常规模型
type Model struct {
	ID        uint      `gorm:"column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

// 软删除处理模型
type SoftModel struct {
	ID        uint           `gorm:"column:id" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`
}

// 微型模型
type TinyModel struct {
	ID        uint      `gorm:"column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
}

// 用户
type User struct {
	SoftModel
	Name         string `gorm:"column:name;size:20"`
	Username     string `gorm:"column:username;size:50;unique"`
	Password     string `gorm:"column:password;size:50"`
	Avatar       string `gorm:"column:avatar;size:100"`
	IsAdmin      bool   `gorm:"column:is_admin"`
	IsGoogleUser bool   `gorm:"column:is_google_user"`
}

// 媒体文件 (实际存储在硬盘中的)
type Media struct {
	TinyModel
	Value     string `gorm:"column:value;size:100;unique"`
	Preview   string `gorm:"column:preview;size:100;unique"`
	Activated bool   `gorm:"column:activated"`
}
