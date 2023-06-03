package models

import (
	"github.com/blackviking27/go-data-management/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Data struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Data{})
}

func (d *Data) CreateData() *Data {
	db.NewRecord(d)
	db.Create(&d)
	return d
}

func GetAllData() []Data {
	var Data []Data
	db.Find(&Data)
	return Data
}

func GetDataById(Id int64) (*Data, *gorm.DB) {
	var getData Data
	db := db.Where("ID=?", Id).Find(&getData)
	return &getData, db
}

func DeleteData(Id int64) Data {
	var data Data
	db.Where("ID=?", Id).Delete(data)
	return data
}
