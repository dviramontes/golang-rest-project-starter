package model

import "github.com/jinzhu/gorm"

type Alarm struct {
	gorm.Model
	// ID int64 | primary key
	Text  string `json: "text"`
	Votes int    `json: "votes"`
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
	if err := db.conn.Order("created_at desc").Find(alarms).Error; err != nil {
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

func (db *DB) Upvote(id int) error {
	var a Alarm
	if err := db.conn.First(&a, id).Error; err != nil {
		return err
	}

	a.Votes = a.Votes + 1
	db.conn.Save(&a)

	return nil
}
