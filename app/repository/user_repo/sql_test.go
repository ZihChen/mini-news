package user_repo

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mini-news/app/model"
	"regexp"
	"testing"
	"time"
)

var (
	mock   sqlmock.Sqlmock
	err    error
	sqlDB  *sql.DB
	gormDB *gorm.DB
)

func TestGetUser(t *testing.T) {
	sqlDB, mock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}

	gormDB, err = gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      sqlDB,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	faker := model.User{
		ID:        1,
		Username:  "jake",
		Password:  "qwe123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	query := "SELECT * FROM `users` WHERE username = ?"

	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs("jake").
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "created_at", "updated_at"}).
			AddRow(faker.ID, faker.Username, faker.Password, faker.CreatedAt, faker.UpdatedAt))

	results := model.User{}
	err = gormDB.Where("username = ?", faker.Username).Find(&results).Error

	assert.Nil(t, err)
	assert.Equal(t, faker, results)
}
