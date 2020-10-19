package models

import (
	"api/security"
	"fmt"
	"time"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json: "id"`
	Nickname  string    `gorm:"size:20;not null;unique" json:"nickname"`
	Email     string    `gorm:"size:50;not null;unique" json:"email"`
	Password  string    `gorm:"size:60;not null" json:"password"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
}

/**
URL: https://gorm.io/docs/hooks.html
*/
func (u *User) BeforeSave() error {
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	fmt.Println(nil)
	return nil
}
