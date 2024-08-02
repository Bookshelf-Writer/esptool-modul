package main

import (
	"esptool/common/cmd"
	"os"
)

//###########################################################//

func (obj *MethodObj) FlashRead() {
	newLog := obj.log.NewLog("FlashRead")
	serialPort := *CLI.Port
	file := *CLI.Flash.FilePath
	size := *CLI.Flash.Size

	{
		if len(serialPort) == 0 {
			newLog.Error().Str("param", CliValPort).Msg(MethodRequiredParameterMissing)
			obj.EndInvalid()
			return
		}
		if len(file) == 0 {
			newLog.Error().Str("param", CliValFlashFile).Msg(MethodRequiredParameterMissing)
			obj.EndInvalid()
			return
		}
		if size == 0 {
			newLog.Error().Str("param", CliValFlashSize).Msg(MethodRequiredParameterMissing)
			obj.EndInvalid()
			return
		}

		_, err := os.Stat(file)
		if err == nil {
			newLog.Error().Msg(MethodFileNotFound)
			obj.EndInvalid()
			return
		}

		if !cmd.IsAccessible(serialPort) {
			newLog.Error().Msg(MethodDevNotAvailable)
			obj.EndInvalid()
			return
		}
	}

	obj.End()
}

func (obj *MethodObj) FlashWrite() {
	newLog := obj.log.NewLog("FlashWrite")
	serialPort := *CLI.Port
	file := *CLI.Flash.FilePath

	{
		if len(serialPort) == 0 {
			newLog.Error().Str("param", CliValPort).Msg(MethodRequiredParameterMissing)
			obj.EndInvalid()
			return
		}
		if len(file) == 0 {
			newLog.Error().Str("param", CliValFlashFile).Msg(MethodRequiredParameterMissing)
			obj.EndInvalid()
			return
		}

		if !cmd.IsAccessible(serialPort) {
			newLog.Error().Msg(MethodDevNotAvailable)
			obj.EndInvalid()
			return
		}
	}

	obj.End()
}
