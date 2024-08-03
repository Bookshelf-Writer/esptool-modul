package output

import "github.com/rs/zerolog"

//###########################################################//

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
