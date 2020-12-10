package models

type User struct {
	Username string `gorm:"size:32"`
	Password []byte `gorm:"size:60"`
}
