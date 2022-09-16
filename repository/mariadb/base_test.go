package mariadb

import (
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestBaseRepo_Updates(t *testing.T) {
	// t.SkipNow()
	dsn := "app_user:app_pass@tcp(127.0.0.1:3306)/app_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	s := ShippingProvider{
		Id: "e8bafded-15fc-480a-b08e-106285878e6f",
	}
	i := Input{Name: "test"}
	err = db.Model(&s).Updates(&i).Error
	if err != nil {
		panic(err)
	}
}

type Input struct {
	Name string `json:"name"`
}

type ShippingProvider struct {
	Id          string `gorm:"type:uuid;primary_key" json:"id"`
	Name        string `json:"name"`
}
