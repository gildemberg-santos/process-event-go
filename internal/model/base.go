package model

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ConectionDB struct {
	DB *gorm.DB
}

func (c *ConectionDB) Open() {
	var err error
	c.DB, err = gorm.Open(sqlite.Open(os.Getenv("DATABASE")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func (d *ConectionDB) Close() {}

func (d *ConectionDB) Migrate(models interface{}) {
	d.DB.AutoMigrate(&models)
}

func (d *ConectionDB) Seed(models interface{}) {
	d.DB.FirstOrCreate(models, models)
}
