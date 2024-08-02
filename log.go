package main

import "github.com/rs/zerolog"

//###########################################################//

const TextLogMsgInit = "INIT"

type LogObj struct {
	log   zerolog.Logger
	index string
}

////

func NewLog(log zerolog.Logger, root string) *LogObj {
	obj := LogObj{index: root}

	newLogger := log.With().Str("index", obj.index).Logger()
	obj.log = newLogger
	newLogger.Debug().Msg(TextLogMsgInit)

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
