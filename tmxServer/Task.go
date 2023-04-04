// Package tmxServer /*
package tmxServer

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"unsafe"
)

type HandleFunctionTypeOfTask func()

type TaskInterface interface {
	Handle() HandleFunctionTypeOfTask
}

type TaskFrameConfig struct {
	//定时器名称
	Name string
	//定时器说明
	Memo string
	//定时任务实例
	Task TaskInterface
	//cron表达式
	Rule string
}

func (t TaskFrameConfig) SetName(taskName string) TaskFrameConfig {
	t.Name = taskName
	return t
}

func (t TaskFrameConfig) SetMemo(taskMemo string) TaskFrameConfig {
	t.Memo = taskMemo
	return t
}

func (t TaskFrameConfig) SetHandleFunc(task TaskInterface) TaskFrameConfig {
	t.Task = task
	return t
}

func (t TaskFrameConfig) SetRule(rule string) TaskFrameConfig {
	t.Rule = rule
	return t
}

type Task struct {
	CornConfig []TaskFrameConfig
	Corn       *cron.Cron
}

func (task *Task) SetConfig(config []TaskFrameConfig) {
	task.CornConfig = config
}

func (task *Task) Key() string {
	return "task"
}

func (task *Task) Handle() *BaseFrame {

	task.Corn = cron.New()

	for _, v := range task.CornConfig {
		func(c TaskFrameConfig, task *Task) {
			_, err := task.Corn.AddFunc(c.Rule, c.Task.Handle())

			if err != nil {
				panic("AddFunc:error:" + err.Error())
			}

			fmt.Println("加载完成")
		}(v, task)
	}

	task.Corn.Start()

	return (*BaseFrame)(unsafe.Pointer(task))
}
