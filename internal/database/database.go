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
	seedData(db)

	DB = db
	fmt.Println("DB接続成功")
}

func seedData(db *gorm.DB) {
	var count int64
	db.Model(&model.Category{}).Count(&count)
	if count > 0 {
		return
	}

	categories := []model.Category{
		{Name: "行政系"},
		{Name: "ライフライン系"},
	}
	db.Create(&categories)

	tasks := []model.Task{
		{CategoryID: categories[0].ID, Title: "転出届", Notes: "引越し前に現住所の市区町村役所へ"},
		{CategoryID: categories[0].ID, Title: "転入届", Notes: "引越し後14日以内に新住所の市区町村役所へ"},
		{CategoryID: categories[0].ID, Title: "運転免許証の住所変更", Notes: "最寄りの警察署または運転免許センターへ"},
		{CategoryID: categories[0].ID, Title: "マイナンバーカードの住所変更", Notes: "転入届と同時に手続き可能"},
		{CategoryID: categories[0].ID, Title: "国民健康保険の住所変更", Notes: "役所で転入届と同時に手続き"},
		{CategoryID: categories[0].ID, Title: "印鑑登録の廃止・新規登録", Notes: "旧住所で廃止、新住所で新規登録"},
		{CategoryID: categories[1].ID, Title: "電気の解約・契約", Notes: "引越し1週間前までに連絡"},
		{CategoryID: categories[1].ID, Title: "ガスの解約・契約", Notes: "立会いが必要なため早めに予約"},
		{CategoryID: categories[1].ID, Title: "水道の解約・契約", Notes: "引越し1週間前までに連絡"},
		{CategoryID: categories[1].ID, Title: "インターネットの移転手続き", Notes: "工事が必要な場合があるため1ヶ月前に連絡"},
		{CategoryID: categories[1].ID, Title: "郵便局への転居届", Notes: "郵便局またはWebで手続き、1年間転送される"},
		{CategoryID: categories[1].ID, Title: "銀行口座の住所変更", Notes: "各銀行のアプリまたは窓口で手続き"},
		{CategoryID: categories[1].ID, Title: "クレジットカードの住所変更", Notes: "各カード会社のWebサイトで手続き"},
		{CategoryID: categories[1].ID, Title: "保険の住所変更", Notes: "各保険会社に連絡"},
	}
	db.Create(&tasks)

	fmt.Println("シードデータ投入完了")
}
