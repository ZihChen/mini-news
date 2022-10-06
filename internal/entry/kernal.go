package entry

import (
	"mini-news/internal/bootstrap"
	"mini-news/internal/schedule"
	"mini-news/internal/server"
)

func Exec() {
	bootstrap.SetupSignal()

	go schedule.Run()

	server.Run()
}
