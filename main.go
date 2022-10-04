package main

import (
	"embed"
	"mini-news/app/global/settings"
	"mini-news/internal/cache"
	"mini-news/internal/database"
	"mini-news/internal/entry"

	_ "mini-news/docs"
)

//go:embed env/*
var f embed.FS

func init() {
	// Loading env
	settings.Load(f)
	// DB connect
	db := database.NewDBInstance()
	db.Ping()
	// Table migrate
	db.CheckTableIsExist()

	// Redis connect
	redis := cache.NewRedisConnect()
	redis.Ping()
}

// @title MiniNews
// @version 1.0
func main() {
	entry.Exec()
}
