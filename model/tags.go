package model

type Tags struct {
	OrderId int    `gorm:"type:int;primary_key"`
	UserId  int    `gorm:"type:int;not null"`
	Name    string `gorm:"type:varchar(255)"`
}
