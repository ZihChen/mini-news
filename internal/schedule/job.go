package schedule

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"mini-news/app/global/errorcode"
	"mini-news/app/global/helper"
	"sync"
	"time"
)

type Job struct {
	Name     string          `json:"name"`      // 背景名稱
	Spec     string          `json:"spec"`      // 執行週期
	FuncName functionName    `json:"func_name"` // 函式名稱
	EntryID  cron.EntryID    `json:"entry_id"`  // EntryID
	wg       *sync.WaitGroup // 等待通道
}

type functionName func() (goErr errorcode.Error)

var Singleton *Job
var Once sync.Once

func NewSeries() *Job {
	Once.Do(func() {
		Singleton = &Job{}
	})
	return Singleton
}

func (j *Job) Init() {
	j.wg = new(sync.WaitGroup)
}

func (j *Job) Run() {
	startTime := time.Now()

	goErr := j.Exec()

	endTime := time.Now()

	j.RecordLog(startTime, endTime, goErr)
}

func (j *Job) LoadSchedule() (jobs []*Job) {

	return []*Job{
		{
			Name: "爬取鏈新聞",
			Spec: "@every 10s",
			FuncName: func() (goErr errorcode.Error) {

				return
			},
		},
	}
}

// Wait 等待WaitGroup結束
func (j *Job) Wait() {
	j.wg.Wait()
}

// Exec 執行function
func (j *Job) Exec() (goErr errorcode.Error) {
	if j.FuncName == nil {
		goErr = helper.ErrorHandle(errorcode.ErrorCronJob, errorcode.FunctionNameNotFound, "")
		return
	}

	return j.FuncName()
}

// RecordLog 紀錄執行時間及錯誤訊息
func (j *Job) RecordLog(startTime, endTime time.Time, goErr errorcode.Error) {
	execTime := endTime.Sub(startTime)
	if goErr != nil {
		fmt.Printf("%v error, error reason %v , and totally spent %v", j.Name, goErr, execTime)
	}

	fmt.Printf("%v execute success, and totally spent %v", j.Name, execTime)
}

func (j *Job) SetEntryID(entryID cron.EntryID) {
	j.EntryID = entryID
}
