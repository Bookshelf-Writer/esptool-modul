package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

//###########################################################//

const (
	TextLogMsgInit = "INIT"
	LvlLogDef      = zerolog.InfoLevel
	LvlLogDebug    = zerolog.DebugLevel
	LvlLogTrase    = zerolog.TraceLevel
)

var (
	LogConsoleColor = log.Output(consoleWriter(false))
	LogConsole      = log.Output(consoleWriter(true))
	LogJson         = zerolog.New(os.Stdout).With().Timestamp().Logger()
)

type LogObj struct {
	log   zerolog.Logger
	index string
}

////

func consoleWriter(NoColor bool) zerolog.ConsoleWriter {
	obj := zerolog.ConsoleWriter{Out: os.Stdout, NoColor: NoColor, TimeFormat: "15:04:05"}

	obj.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%-16s", i)
	}

	obj.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%-6s", i)
	}
	return obj
}

func NewLog(log zerolog.Logger, root string) *LogObj {
	obj := LogObj{index: root}

	newLogger := log.With().Str("index", obj.index).Logger()
	obj.log = newLogger

	return &obj
}

func (obj *LogObj) NewLog(point string) *LogObj {
	newObj := LogObj{index: obj.index + "/" + point}

	newLogger := obj.log.With().Str("index", newObj.index).Logger()
	newObj.log = newLogger
	newLogger.Debug().Msg(TextLogMsgInit)

	return &newObj
}

////

func (obj *LogObj) Trace() *zerolog.Event {
	return obj.log.Trace()
}
func (obj *LogObj) Debug() *zerolog.Event {
	return obj.log.Debug()
}
func (obj *LogObj) Info() *zerolog.Event {
	return obj.log.Info()
}
func (obj *LogObj) Warn() *zerolog.Event {
	return obj.log.Warn()
}
func (obj *LogObj) Error() *zerolog.Event {
	return obj.log.Error()
}
func (obj *LogObj) Fatal() *zerolog.Event {
	return obj.log.Fatal()
}
func (obj *LogObj) Panic() *zerolog.Event {
	return obj.log.Panic()
}

////

type StringArray []string

func (a StringArray) MarshalZerologArray(arr *zerolog.Array) {
	for _, s := range a {
		arr.Str(s)
	}
}
