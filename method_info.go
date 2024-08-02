package main

import "esptool/common/cmd"

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

	esp32, err := connectEsp32(serialPort, uint32(*CLI.Baud.Connect), uint32(*CLI.Baud.Transfer), *CLI.Conn.Retries, newLog)
	if err != nil {
		newLog.Error().Err(err).Msg("connect esp32 failed")
		obj.EndInvalid()
		return
	}

	err = infoCommand(esp32)
	newLog.Error().Err(err).Send()

	obj.End()
}
