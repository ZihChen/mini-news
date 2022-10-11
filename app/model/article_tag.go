package model

import "time"

type ArticleTag struct {
	ID        int       `json:"id" gorm:"column:id;type:int(11) unsigned auto_increment;NOT NULL;primary_key"`
	ArticleID int       `json:"article_id" gorm:"column:article_id;type:int(11);NOT NULL"`
	TagID     int       `json:"tag_id" gorm:"column:tag_id;type:int(11);NOT NULL"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:TIMESTAMP; default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:TIMESTAMP; not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
