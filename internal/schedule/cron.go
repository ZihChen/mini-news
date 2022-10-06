package schedule

import (
	"fmt"
	"github.com/robfig/cron/v3"
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
	fmt.Printf("[⛑] CronJob start to run!")
}
