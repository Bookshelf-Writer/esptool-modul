package output

import "github.com/rs/zerolog"

//###########################################################//

type LogObj struct {
	log   zerolog.Logger
	index string
}

////

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

func (obj *LogObj) ZeroLog() *zerolog.Logger {
	return &obj.log
}
