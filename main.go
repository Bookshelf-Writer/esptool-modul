package main

import (
	"flag"
	"strings"
)

func main() {
	logLVL := LvlLogDef
	bufLog := LogConsoleColor
	flag.Parse()

	//manual rechecking of parameters
	for _, arg := range flag.Args() {
		switch strings.ToLower(arg) {
		case strings.ToLower(CliTrigHelp), "h":
			*CliTrigType.Help = true

		case "logtrase":
			logLVL = LvlLogTrase
		case "logdebug":
			logLVL = LvlLogDebug

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
		bufLog = LogConsole
	case *CliTrigType.Json:
		bufLog = LogJson
	}

	logs := NewLog(bufLog.Level(logLVL), "CLI")
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
