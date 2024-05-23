package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Thiti-Dev/ksc/internal/cmds"
	"github.com/common-nighthawk/go-figure"
)

func main() {
	cmds.AddAllCommandsTo(cmds.RootCmd) // sub commands registration
	figure.NewFigure("KSC", "", true).Print()
	time.Sleep(time.Millisecond * 500) // 500 milisecond until loading up the tui

	if err := cmds.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
