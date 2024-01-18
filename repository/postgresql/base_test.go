package postgresql

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func TestBaseRepo_Updates(t *testing.T) {
	t.SkipNow()
	dsn := "user=postgres password=qweasd DB.name=wallets_db host=localhost port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DryRun: true,
	})
	db = db.Debug()
	if err != nil {
		panic(err)
	}
	s := ShippingProvider{
		Id: "e8bafded-15fc-480a-b08e-106285878e6f",
	}
	i := Input{Name: "test"}
	db.Model(&s).Updates(&i)
	baseRepo := NewBaseRepo(db)
}

type Input struct {
	Name string `json:"name"`
}

type ShippingProvider struct {
	Id        string `gorm:"type:uuid;primary_key" json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
	Active    *bool  `json:"active"`
	Slug      string `json:"slug"`
	AddressID string `json:"addressID"`
	LogoImage string `json:"logoImage"`
	CreatedBy string `json:"createdBy"`
	UpdatedBy string `json:"updatedBy"`
}
