package schedule

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"mini-news/internal/bootstrap"
	"os"
)

func Run() {
	series := NewSeries()
	jobs := series.LoadSchedule()

	c := cron.New(cron.WithSeconds())

	for _, job := range jobs {

		entryID, err := c.AddJob(job.Spec, job)
		if err != nil {

		}

		job.SetEntryID(entryID)
	}

	c.Start()
	fmt.Printf("⛑ 啟動排程 ⛑")
	<-bootstrap.GracefulDown()

	select {
	case <-bootstrap.WaitOnceSignal():
		fmt.Println("🚦  收到關閉訊號，強制結束 🚦")

		// 等待背景結束
		for _, job := range jobs {
			fmt.Println(job)
			job.Wait()
		}
		os.Exit(2)

	}
}
