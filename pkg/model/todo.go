package model

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Text string
	Completed bool
}

func (db *DB) GetAllTodos(todos *[]Todo) error {
	if err := db.conn.Find(todos).Error; err != nil {
		return err
	}
	return nil
}

func (db *DB) CreateTodo(todo *Todo) error {
	if err := db.conn.Create(todo).Error; err != nil {
		return err
	}
	return  nil
}
