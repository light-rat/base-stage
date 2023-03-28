package model

type


User struct {
	Id    int `gorm:"primary_key"`
	Name  string
	Home  string
	Email string
	Sex   int
	TableBase
}
