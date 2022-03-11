package models

import (
	"github.com/jinzhu/gorm"
	"github.com/rrscodes/fashionfever-api/pkg/config"
)

var db *gorm.DB

type Dress struct {
	gorm.Model
	Label string `gorm:""json:"label"`
	Brand string `json:"brand"`
	Size  string `json:"size"`
	Color string `json:"color"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Dress{})
}

func (d *Dress) CreateDress() *Dress {
	db.NewRecord(d)
	db.Create(&d)
	return d
}

func GetAllDresses() []Dress {
	var Dresses []Dress
	db.Find(&Dresses)
	return Dresses
}

func GetDressById(Id int64) (*Dress, *gorm.DB) {
	var getDress Dress
	db := db.Where("id = ?", Id).Find(&getDress)
	return &getDress, db
}

func DeleteDress(Id int64) Dress {
	var dress Dress
	db.Where("id = ?", Id).Delete(dress)
	return dress
}
