package model

type Todo struct {
	ID     int64  `json:"id" gorm:"column:id"`
	Name   string `json:"name" gorm:"column:name"`
	Status bool   `json:"status" gorm:"column:status"`
}

func (t Todo) TableName() string {
	return "todo_list"
}
