package main

import (
	"flag"
	"github.com/Bookshelf-Writer/esptool-modul/lib/output"
	"strings"
)

func main() {
	logLVL := output.LvlLogDef
	bufLog := output.LogConsoleColor
	flag.Parse()

	//manual rechecking of parameters
	for _, arg := range flag.Args() {
		switch strings.ToLower(arg) {
		case strings.ToLower(CliTrigHelp), "h":
			*CliTrigType.Help = true

		case "logtrase":
			logLVL = output.LvlLogTrase
		case "logdebug":
			logLVL = output.LvlLogDebug

		case "v", "ver":
			*CliTrigType.Version = true

		default:
			fl, ok := CliTrigTypeMap[arg]
			if ok {
				*fl = true
			}
		}
	}

	//preliminary interception of conditions
	switch {
	case *CliTrigType.NoColor:
		bufLog = output.LogConsole
	case *CliTrigType.Json:
		bufLog = output.LogJson
	}

	logs := output.NewLog(bufLog.Level(logLVL), "CLI")
	method := MethodObj{log: logs}

	//move by group
	switch {
	case *CliTrigType.Version:
		method.Version()
		return

	case *CliTrigType.List:
		method.List()
		return

	case *CliTrigType.Info:
		method.Info()
		return

	case *CliTrigType.FlashRead:
		method.FlashRead()
		return

	case *CliTrigType.FlashWrite:
		method.FlashWrite()
		return
	}

	method.Help()
}
