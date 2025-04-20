package handler

import (
	"time"

	"gorm.io/gorm"
)

type Page struct {
	Total   int64 `json:"total"`
	Page    int64 `json:"page"`
	Content any   `json:"content"`
}

type Model struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type SoftModel struct {
	ID        uint           `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
