// Package tmxServer /*
package tmxServer

import (
	"fmt"
	"time"
	"unsafe"
)

type ProcessConfig struct {
	//process名称
	ProcessName string
	//当前process工作数量
	ProcessNumber int
	//process运行实例
	ProcessFunc ProcessInterface
}

type ProcessInterface interface {
	ProcessHandle(entry ProcessEntry)
}

// ProcessEntry
// @Description: process当前工作协程空间信息
type ProcessEntry struct {
	//第几个协程
	Index int
	//process 名称
	Name string
}

// Process
// @Description: 工作process协程组信息
type Process struct {
	ProcessConfig []ProcessConfig
	entryList     []ProcessEntry
}

func (process *Process) SetConfig(pC []ProcessConfig) {
	process.ProcessConfig = pC
}

func (process *Process) Key() string {
	return "process"
}

func (process *Process) Handle() *BaseFrame {
	process.entryList = make([]ProcessEntry, 0)

	for _, config := range process.ProcessConfig {
		func(c ProcessConfig, process *Process) {
			for i := 1; i <= c.ProcessNumber; i++ {
				entry := ProcessEntry{
					Index: i,
					Name:  c.ProcessName,
				}

				process.entryList = append(process.entryList, entry)

				go process.entryHandle(c, entry)
			}
		}(config, process)
	}

	return (*BaseFrame)(unsafe.Pointer(process))
}

func (process *Process) entryHandle(p ProcessConfig, e ProcessEntry) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("process发生异常", err)
		}

		time.Sleep(2 * time.Second)

		process.entryHandle(p, e)
	}()

	p.ProcessFunc.ProcessHandle(e)
}
