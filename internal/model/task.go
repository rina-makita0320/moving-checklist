package model

import "time"

type Task struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	CategoryID uint       `json:"category_id"`
	Category   *Category  `json:"category"`
	Title      string     `gorm:"not null" json:"title"`
	DueDate    *time.Time `json:"due_date"`
	Completed  bool       `gorm:"default:false" json:"completed"`
	Notes      string     `json:"notes"`
	CreatedAt  time.Time  `json:"created_at"`
}

type Category struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"not null" json:"name"`
	Tasks []Task `gorm:"foreignKey:CategoryID" json:"-"`
}
