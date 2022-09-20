package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"mini-news/app/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var r *gin.Engine

func Run() {

	r = gin.New()

	router.LoadingRouter(r)

	srv := http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("[❌ Fatal❌ ] Server 建立監聽連線失敗:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("🚦 收到訊號囉，等待其他連線完成，準備結束服務 🚦")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("🚦 收到關閉訊號，強制結束 🚦")
	}
}
