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

	/*
		esp32, err := connectEsp32(*flashReadPort, uint32(*flashReadConnectBaudrate), uint32(*flashReadTransferBaudrate), *flashReadRetries, logger)
		if err != nil {
			return err
		}
		bytes, err := esp32.ReadFlash(uint32(*flashReadOffset), uint32(*flashReadSize))
		if err != nil {
			return err
		}
		os.Stdout.Write(bytes)
		return nil
	*/

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

	/*
		contents, err := ioutil.ReadFile(*flashWriteFile)
		if err != nil {
			return err
		}
		esp32, err := connectEsp32(*flashWritePort, uint32(*flashWriteConnectBaudrate), uint32(*flashWriteTransferBaudrate), *flashWriteRetries, logger)
		if err != nil {
			return err
		}

		err = esp32.WriteFlash(uint32(*flashWriteOffset), contents, *flashWriteCompress)
		if err != nil {
			panic(err)
		}
		logger.Print("Done")
		return nil
	*/

	obj.End()
}
