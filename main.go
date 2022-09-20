package main

import (
	"mini-news/app/global/config"
	"mini-news/internal/entry"
)

func init() {
	config.Load()
}

func main() {
	entry.Exec()
}
