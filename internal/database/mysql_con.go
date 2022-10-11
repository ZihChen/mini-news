package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"mini-news/app/global/errorcode"
	"mini-news/app/global/helper"
	"mini-news/app/global/settings"
	"mini-news/app/model"
	"time"
)

var connectPool *gorm.DB

type Interface interface {
	DBConnect() (*gorm.DB, errorcode.Error)
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

func (d *dbCon) DBConnect() (*gorm.DB, errorcode.Error) {
	var err error

	if connectPool != nil {
		return connectPool, nil
	}

	dsn := composeConfString()

	connectPool, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		goErr := helper.ErrorHandle(errorcode.ErrorDatabase, errorcode.DBConnectError, err.Error())
		return nil, goErr
	}

	sqlPool, err := connectPool.DB()
	if err != nil {
		goErr := helper.ErrorHandle(errorcode.ErrorDatabase, errorcode.DBConnectPoolError, err.Error())
		return nil, goErr
	}

	// 限制最大開啟的連線數
	sqlPool.SetMaxIdleConns(100)
	// 限制最大閒置連線數
	sqlPool.SetMaxOpenConns(2000)
	// 空閒連線 timeout 時間
	sqlPool.SetConnMaxLifetime(15 * time.Second)

	return connectPool.Debug(), nil
}

func (d *dbCon) Ping() {
	connectPool, goErr := d.DBConnect()
	if goErr != nil {
		log.Fatalf(errorcode.CheckDBConnectError, goErr.GetErrorMsg())
	}

	db, err := connectPool.DB()
	if err != nil {
		log.Fatalf(errorcode.CheckConnectPoolError, err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf(errorcode.PingDBError, err.Error())
	}
}

func (d *dbCon) CheckTableIsExist() {
	db, goErr := d.DBConnect()
	if goErr != nil {
		log.Fatalf(errorcode.CheckDBConnectError, goErr.GetErrorMsg())
	}

	err := db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf(errorcode.UserTableMigrateError, err.Error())
	}

	err = db.AutoMigrate(&model.CryptoArticle{})
	if err != nil {
		log.Fatalf(errorcode.CryptoArticleTableMigrateError, err.Error())
	}

	err = db.AutoMigrate(&model.Tag{})
	if err != nil {
		log.Fatalf(errorcode.TagTableMigrateError, err.Error())
	}

	err = db.AutoMigrate(&model.ArticleTag{})
	if err != nil {
		log.Fatalf(errorcode.ArticleTagTableMigrateError, err.Error())
	}
}

func composeConfString() string {
	Host := settings.Config.DBConfig.Host
	Username := settings.Config.DBConfig.Username
	Password := settings.Config.DBConfig.Password
	Database := settings.Config.DBConfig.Database

	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?timeout=5s&charset=utf8mb4&parseTime=True&loc=Local", Username, Password, Host, Database)
}
