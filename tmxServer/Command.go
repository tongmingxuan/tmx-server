// Package tmxServer /*
package tmxServer

import (
	"github.com/urfave/cli"
	"os"
	"unsafe"
)

// CommandItem
// @Description: 自定义command实例
type CommandItem interface {
	Command() cli.Command
}

type Command struct {
	CommandList []cli.Command
	CliApp      *cli.App
}

// AppendCommand
// @Description: 设置自定义command集合
// @receiver command
// @param commands
func (command *Command) AppendCommand(commands ...CommandItem) {
	var cliCommands []cli.Command

	for _, item := range commands {
		cliCommands = append(cliCommands, item.Command())
	}

	command.CommandList = cliCommands
}

func (command *Command) SetConfig(config []CommandItem) {
	command.AppendCommand(config...)
}

func (command *Command) Key() string {
	return "command"
}

func (command *Command) Handle() *BaseFrame {
	cliApp := cli.NewApp()
	cliApp.Name = "tmx-server"
	cliApp.Usage = "tmx-server"
	cliApp.Email = "1422906354@qq.com"
	cliApp.Version = "2.0.0"
	cliApp.Usage = "自定义command"

	cliApp.Commands = command.CommandList

	err := cliApp.Run(os.Args)

	if err != nil {
		panic("cliApp:error:" + err.Error())
	}

	command.CliApp = cliApp

	return (*BaseFrame)(unsafe.Pointer(command))
}
