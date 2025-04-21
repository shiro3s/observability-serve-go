package dao

import (
	"fmt"
	"template_app/models"

	"gorm.io/gorm"
)

func SearchTodo(db *gorm.DB, cond models.TodoSearchParam) ([]models.Todo, int64) {
	session := db.Table("t_todos")
	todos := []models.Todo{}
	var count int64

	if cond.Done {
		session.Where("done = ?", true)
	}

	if cond.Title != "" {
		t := "%" + cond.Title + "%"
		session.Where("title LIKE ?", t)
	}

	if cond.Limit != 0 {
		session.Limit(cond.Limit)
	}

	if cond.Offset != 0 {
		session.Offset(cond.Offset)
	}
	order := fmt.Sprintf("%s %s", cond.Sort, cond.Order)
	session.Order(order)

	session.Where("is_deleted = ?", cond.IsDeleted).Find(&todos).Count(&count)

	return todos, count
}

func FindTodoById(db *gorm.DB, id int, cond models.TodoParam) (*models.Todo, error) {
	session := db.Table("t_todos")
	todo := models.Todo{}

	session.Where("id = ? and is_deleted = ?", id, true)
	result := session.Take(&todo)

	if result.Error != nil {
		fmt.Println("FindTodoByID SQL ERROR:", result.Error)
		return nil, result.Error
	}

	return &todo, nil
}

func CreateTodo(db *gorm.DB, postBody models.TodoBody) error {
	session := db.Table("t_todos")
	result := session.Create(&postBody)

	if result.Error != nil {
		fmt.Println("CreateTodo SQL ERROR:", result.Error)
		return result.Error
	}

	return nil
}

func UpdateTodo(db *gorm.DB, id int, putBody models.TodoBody) error {
	session := db.Table("t_todos")

	result := session.Where("id = ?", id).Updates(&putBody)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteTodo(db *gorm.DB, id int) error {
	session := db.Table("t_todos")

	result := session.Where("id = ?", id).Update("is_deleted", true)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
