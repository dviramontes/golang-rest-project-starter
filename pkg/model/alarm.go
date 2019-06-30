package model

import (
	"log"

	"github.com/jinzhu/gorm"
)

type DB struct {
	conn *gorm.DB
}

func New(db *gorm.DB) *DB {
	return &DB{db}
}

type Alarm struct {
	gorm.Model
	// ID int64 | primary key
	Text string
}

func (db *DB) Migrate() {
	db.conn.AutoMigrate(&Alarm{})
}

func (db *DB) Seed() {
	log.Println("seeding db...")

	db.CreateAlarm("This is the first alarm.")
	db.CreateAlarm("This is the second alarm.")
}

func (db *DB) CreateAlarm(text string) {
	db.conn.Create(&Alarm{Text: text})
}

func (db *DB) GetAllAlarms(alarms *[]Alarm) {
	db.conn.Find(alarms)
}

func (db *DB) DeleteAlarm(id int64) {
	db.conn.Unscoped().Delete(&Alarm{}, "id = ?", id)
}

func (db *DB) DeleteAllAlarms() {
	db.conn.Unscoped().Delete(&Alarm{})
}
