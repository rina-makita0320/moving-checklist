package model

import "time"

type Task struct {
	ID         uint   `gorm:"primaryKey"`
	CategoryID uint
	Title      string `gorm:"not null"`
	DueDate    *time.Time
	Completed  bool   `gorm:"default:false"`
	Notes      string
	CreatedAt  time.Time
}

type Category struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Tasks []Task `gorm:"foreignKey:CategoryID"`
}