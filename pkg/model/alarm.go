package model

import "github.com/jinzhu/gorm"

type Alarm struct {
	gorm.Model
	// ID int64 | primary key
	Text string `json: "text"`
}

func (db *DB) Seed() {
	a := &Alarm{Text: "This is the first alarm."}
	b := &Alarm{Text: "This is the second alarm."}
	db.conn.Where(Alarm{Text: "This is the first alarm."}).FirstOrCreate(a)
	db.conn.Where(Alarm{Text: "This is the second alarm."}).FirstOrCreate(b)
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

func (db *DB) DeleteAlarm(id int) error {
	if err := db.conn.Delete(&Alarm{}, "id = ?", id).Error; err != nil {
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
