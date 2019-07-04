package model

import (
	"github.com/jinzhu/gorm"
)

type DB struct {
	conn *gorm.DB
}

func New(db *gorm.DB) *DB {
	return &DB{db}
}

func (db *DB) Migrate() {
	db.conn.AutoMigrate(&Alarm{}, &Todo{})
}
