package database

import (
	"fmt"
	"log"

	"github.com/rina-makita0320/moving-checklist/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// ローカル開発用のため sslmode=disable、本番環境では要変更
	dsn := "host=localhost user=postgres password=password dbname=moving_checklist port=5432 sslmode=disable TimeZone=Asia/Tokyo client_encoding=UTF8"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DBに接続できませんでした:", err)
	}

	db.AutoMigrate(&model.Category{}, &model.Task{})

	DB = db
	fmt.Println("DB接続成功")
}
