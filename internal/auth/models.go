package auth

import "time"

type User struct {
	ID       uint   `gorm:"primaryKey,autoIncrement,index"`
	Name     string `gorm:"unique"`
	Password string
	Created  time.Time `gorm:"autoCreateTime"`
}
