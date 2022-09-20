package entry

import (
	"mini-news/internal/bootstrap"
	"mini-news/internal/server"
)

func Exec() {
	bootstrap.SetupSignal()

	server.Run()
}