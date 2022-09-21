package model

import "time"

type User struct {
	ID        int       `json:"id" gorm:"column:id;type:int(11) unsigned auto_increment;NOT NULL;primary_key"`
	Username  string    `json:"username" gorm:"column:username;type:varchar(255);not null"`
	Password  string    `json:"password" gorm:"column:password;type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:TIMESTAMP; default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:TIMESTAMP; not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
