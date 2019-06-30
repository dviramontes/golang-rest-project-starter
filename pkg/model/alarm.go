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
	a := &Alarm{Text: "This is the first alarm."}
	b := &Alarm{Text: "This is the second alarm."}
	db.CreateAlarm(a)
	db.CreateAlarm(b)
}

func (db *DB) CreateAlarm(alarm *Alarm) error {
	if err := db.conn.Create(alarm).Error; err != nil {
		return err
	}
	return nil
}

func (db *DB) GetAllAlarms(alarms *[]Alarm) error {
	if err := db.conn.Find(alarms).Error; err != nil {
		return err
	}
	return nil
}

func (db *DB) DeleteAlarm(id int64) error {
	if err := db.conn.Unscoped().Delete(&Alarm{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (db *DB) DeleteAllAlarms() error {
	if err := db.conn.Unscoped().Delete(&Alarm{}).Error; err != nil {
		return err
	}
	return nil
}
