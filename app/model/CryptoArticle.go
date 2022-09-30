package model

import "time"

type CryptoArticle struct {
	ID        int       `json:"id" gorm:"column:id;type:int(11) unsigned auto_increment;NOT NULL;primary_key"`
	Name      string    `json:"name" gorm:"column:name;type:varchar(255) comment '文章名稱';not null"`
	Title     string    `json:"title" gorm:"column:title;type:varchar(127) comment '文章標題';not null"`
	Site      string    `json:"site" gorm:"column:site;type:varchar(127) comment '文章來源';not null"`
	Source    string    `json:"source" gorm:"column:source;type:varchar(255) comment '文章連結';not null"`
	Image     string    `json:"image" gorm:"column:image;type:varchar(255) comment '文章圖片連結';not null"`
	Date      time.Time `json:"date" gorm:"column:date;type:datetime comment '遊戲點擊紀錄日期';not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:TIMESTAMP; default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:TIMESTAMP; not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
