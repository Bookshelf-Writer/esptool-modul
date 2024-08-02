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

	obj.End()
}
