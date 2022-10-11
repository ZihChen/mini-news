package model

import "time"

type Tag struct {
	ID        int       `json:"id" gorm:"column:id;type:int(11) unsigned auto_increment;NOT NULL;primary_key"`
	Name      string    `json:"name" gorm:"column:name;type:varchar(255);not null"`
	UserID    int       `json:"user_id" gorm:"column:user_id;type:int(11);NOT NULL"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:TIMESTAMP; default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:TIMESTAMP; not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
