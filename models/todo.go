package models

import "time"

type Todo struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Done      *bool     `json:"done"`
	IsDeleted *bool     `json:"is_deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TodoSearchParam struct {
	Title     string    `query:"title"`
	Done      bool      `query:"done"`
	IsDeleted bool      `query:"is_deleted"`
	Sort      string    `query:"sort"`
	Order     OrderEnum `query:"order"`
	Limit     int       `query:"limit"`
	Offset    int       `query:"offset"`
}

type TodoParam struct {
	IsDeleted bool `query:"is_deleted"`
}

type TodoBody struct {
	Title string `form:"title"`
	Done  *bool  `form:"done" gorm:"default:false"`
}

func NewTodoSearchParameter() TodoSearchParam {
	return TodoSearchParam{
		IsDeleted: false,
		Sort:      "id",
		Order:     OrderDesc,
		Offset:    0,
		Limit:     30,
	}
}
