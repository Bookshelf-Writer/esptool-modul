package main

import (
	"github.com/Bookshelf-Writer/esptool-modul"
	"github.com/Bookshelf-Writer/esptool-modul/common/cmd"
)

//###########################################################//

func (obj *MethodObj) Info() {
	newLog := obj.log.NewLog("Info")
	serialPort := *CLI.Port

	if len(serialPort) == 0 {
		newLog.Error().Str("param", CliValPort).Msg(MethodRequiredParameterMissing)
		obj.EndInvalid()
		return
	}

	if !cmd.IsAccessible(serialPort) {
		newLog.Error().Msg(MethodDevNotAvailable)
		obj.EndInvalid()
		return
	}

	//todo причесать нормально
	{
		esp32, err := esptool.ConnectEsp32(serialPort, uint32(*CLI.Baud.Connect), uint32(*CLI.Baud.Transfer), *CLI.Conn.Retries, newLog)
		if err != nil {
			newLog.Error().Err(err).Msg("connect esp32 failed")
			obj.EndInvalid()
			return
		}

		err = esptool.InfoCommand(esp32)
		if err != nil {
			newLog.Error().Err(err).Msg("esp32 info failed")
			obj.EndInvalid()
			return
		}

		newLog.Info().Msg("OK")
	}

	obj.End()
}
