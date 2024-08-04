package main

import (
	"github.com/Bookshelf-Writer/esptool-modul"
	cmd "github.com/Bookshelf-Writer/esptool-modul/lib/serial"
	"io/ioutil"
	"os"
)

//###########################################################//

func (obj *MethodObj) FlashRead() {
	newLog := obj.log.NewLog("FlashRead")
	serialPort := *CLI.Port
	filename := *CLI.Flash.FilePath
	size := *CLI.Flash.Size

	{
		if len(serialPort) == 0 {
			newLog.Error().Str("param", CliValPort).Msg(MethodRequiredParameterMissing)
			obj.EndInvalid()
			return
		}
		if len(filename) == 0 {
			newLog.Error().Str("param", CliValFlashFile).Msg(MethodRequiredParameterMissing)
			obj.EndInvalid()
			return
		}
		if size == 0 {
			newLog.Error().Str("param", CliValFlashSize).Msg(MethodRequiredParameterMissing)
			obj.EndInvalid()
			return
		}

		if !cmd.Check(serialPort) {
			newLog.Error().Msg(MethodDevNotAvailable)
			obj.EndInvalid()
			return
		}
	}

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		newLog.Error().Err(err).Msg(MethodFileNotFound)
		obj.EndInvalid()
		return
	}
	defer file.Close()

	//todo причесать нормально
	{
		esp32, err := esptool.ConnectEsp32(serialPort, uint32(*CLI.Baud.Connect), uint32(*CLI.Baud.Transfer), *CLI.Conn.Retries, newLog)
		if err != nil {
			newLog.Error().Err(err).Msg("connect esp32 failed")
			obj.EndInvalid()
			return
		}
		bytes, err := esp32.ReadFlash(uint32(*CLI.Flash.Offset), uint32(size))
		if err != nil {
			newLog.Error().Err(err).Msg("read esp32 failed")
			obj.EndInvalid()
			return
		}

		_, err = file.Write(bytes)
		if err != nil {
			newLog.Error().Err(err).Msg("write esp32 failed")
			obj.EndInvalid()
			return
		}

		newLog.Info().Msg("OK")
	}
	obj.End()
}

////////

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

		if !cmd.Check(serialPort) {
			newLog.Error().Msg(MethodDevNotAvailable)
			obj.EndInvalid()
			return
		}
	}

	//todo причесать нормально
	{
		contents, err := ioutil.ReadFile(file)
		if err != nil {
			newLog.Error().Err(err).Msg("error when trying to create a file")
			obj.EndInvalid()
			return
		}
		esp32, err := esptool.ConnectEsp32(serialPort, uint32(*CLI.Baud.Connect), uint32(*CLI.Baud.Transfer), *CLI.Conn.Retries, newLog)
		if err != nil {
			newLog.Error().Err(err).Msg("connect esp32 failed")
			obj.EndInvalid()
			return
		}

		newLog.Debug().Int("size", len(contents)).Msg("file")
		err = esp32.WriteFlash(uint32(*CLI.Flash.Offset), contents, *CLI.Flash.Compress)
		if err != nil {
			newLog.Error().Err(err).Msg("write esp32 failed")
			obj.EndInvalid()
			return
		}

		newLog.Info().Msg("OK")
	}

	obj.End()
}
