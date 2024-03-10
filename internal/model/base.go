package model

import (
	"crypto/sha1"
	"crypto/sha512"
	"encoding/hex"
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

func ParseStringToSha512(str string) string {
	hasher := sha512.New()
	hasher.Write([]byte(str))
	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}

func ParseStringToSha1(str string) string {
	hasher := sha1.New()
	hasher.Write([]byte(str))
	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}
