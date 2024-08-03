package output

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
)

//###########################################################//

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
