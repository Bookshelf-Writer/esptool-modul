package output

import (
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
