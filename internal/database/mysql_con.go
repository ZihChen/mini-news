package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"mini-news/app/global/settings"
	"mini-news/app/model"
	"time"
)

var connectPool *gorm.DB

type Interface interface {
	DBConnect() (*gorm.DB, error)
	CheckTableIsExist()
	Ping()
}

type dbCon struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func NewDBInstance() Interface {
	return &dbCon{}
}

func (d *dbCon) DBConnect() (*gorm.DB, error) {
	var err error

	if connectPool != nil {
		return connectPool, nil
	}

	dsn := composeConfString()

	connectPool, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf("DB CONNECT ERROR: %v", err.Error())
	}

	sqlPool, err := connectPool.DB()
	if err != nil {
		log.Fatalf("DB CONNECT POOL ERROR: %v", err.Error())
	}

	// 限制最大開啟的連線數
	sqlPool.SetMaxIdleConns(100)
	// 限制最大閒置連線數
	sqlPool.SetMaxOpenConns(2000)
	// 空閒連線 timeout 時間
	sqlPool.SetConnMaxLifetime(15 * time.Second)

	return connectPool.Debug(), err
}

func (d *dbCon) Ping() {
	connectPool, err := d.DBConnect()
	if err != nil {
		log.Fatalf("CHECK DB CONNECT ERROR: %v", err.Error())
	}

	db, err := connectPool.DB()
	if err != nil {
		log.Fatalf("CHECK CONNECT POOL ERROR: %v", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("PING DB ERROR: %v", err.Error())
	}
}

func (d *dbCon) CheckTableIsExist() {
	db, err := d.DBConnect()
	if err != nil {
		log.Fatalf("CHECK DB CONNECT ERROR: %v", err.Error())
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("USER TABLE MIGRATE ERROR: %v", err.Error())
	}

	err = db.AutoMigrate(&model.CryptoArticle{})
	if err != nil {
		log.Fatalf("CREYTOARTICLE TABLE MIGRATE ERROR: %v", err.Error())
	}
}

func composeConfString() string {
	Host := settings.Config.DBConfig.Host
	Username := settings.Config.DBConfig.Username
	Password := settings.Config.DBConfig.Password
	Database := settings.Config.DBConfig.Database

	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?timeout=5s&charset=utf8mb4&parseTime=True&loc=Local", Username, Password, Host, Database)
}
