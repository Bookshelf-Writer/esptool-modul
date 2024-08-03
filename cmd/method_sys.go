package main

import (
	"fmt"
	"github.com/Bookshelf-Writer/esptool-modul"
	"github.com/Bookshelf-Writer/esptool-modul/common/cmd"
)

//###########################################################//

func (obj *MethodObj) Help() {
	newLog := obj.log.NewLog("Help")

	newLog.Warn().Str("usage", CliTrigTextVersion).Msg(CliTrigVersion)
	newLog.Info().Str("usage", CliTrigTextJson).Msg(CliTrigJson)
	newLog.Info().Str("usage", CliTrigTextNoColor).Msg(CliTrigNoColor)
	fmt.Println()

	//

	listLog := newLog.NewLog("List")
	listLog.Warn().Str("usage", CliTrigTextList).Msg(CliTrigList)
	fmt.Println()

	//

	infoLog := newLog.NewLog("Info")
	infoLog.Warn().Str("usage", CliTrigTextInfo).Msg(CliTrigInfo)
	infoLog.Info().Str("usage", CliValTextPort).Bool("req", true).Msg(CliValPort)
	infoLog.Info().Str("usage", CliValTextBaudConnect).Bool("req", false).Msg(CliValBaudConnect)
	infoLog.Info().Str("usage", CliValTextBaudTransfer).Bool("req", false).Msg(CliValBaudTransfer)
	infoLog.Info().Str("usage", CliValTextConnTimeout).Bool("req", false).Msg(CliValConnTimeout)
	infoLog.Info().Str("usage", CliValTextConnRetries).Bool("req", false).Msg(CliValConnRetries)
	fmt.Println()

	//

	flashReadLog := newLog.NewLog("FlashRead")
	flashReadLog.Warn().Str("usage", CliTrigTextFlashRead).Msg(CliTrigFlashRead)
	flashReadLog.Info().Str("usage", CliValTextPort).Bool("req", true).Msg(CliValPort)
	flashReadLog.Info().Str("usage", CliValTextBaudConnect).Bool("req", false).Msg(CliValBaudConnect)
	flashReadLog.Info().Str("usage", CliValTextBaudTransfer).Bool("req", false).Msg(CliValBaudTransfer)
	flashReadLog.Info().Str("usage", CliValTextConnTimeout).Bool("req", false).Msg(CliValConnTimeout)
	flashReadLog.Info().Str("usage", CliValTextConnRetries).Bool("req", false).Msg(CliValConnRetries)

	flashReadLog.Info().Str("usage", CliValTextFlashFile).Bool("req", true).Msg(CliValFlashFile)
	flashReadLog.Info().Str("usage", CliValTextFlashSize).Bool("req", true).Msg(CliValFlashSize)
	flashReadLog.Info().Str("usage", CliValTextFlashOffset).Bool("req", false).Msg(CliValFlashOffset)
	fmt.Println()

	//

	flashWriteLog := newLog.NewLog("FlashWrite")
	flashWriteLog.Warn().Str("usage", CliTrigTextFlashWrite).Msg(CliTrigFlashWrite)
	flashWriteLog.Info().Str("usage", CliValTextPort).Bool("req", true).Msg(CliValPort)
	flashWriteLog.Info().Str("usage", CliValTextBaudConnect).Bool("req", false).Msg(CliValBaudConnect)
	flashWriteLog.Info().Str("usage", CliValTextBaudTransfer).Bool("req", false).Msg(CliValBaudTransfer)
	flashWriteLog.Info().Str("usage", CliValTextConnTimeout).Bool("req", false).Msg(CliValConnTimeout)
	flashWriteLog.Info().Str("usage", CliValTextConnRetries).Bool("req", false).Msg(CliValConnRetries)

	flashWriteLog.Info().Str("usage", CliValTextFlashFile).Bool("req", true).Msg(CliValFlashFile)
	flashWriteLog.Info().Str("usage", CliValTextFlashSize).Bool("req", false).Msg(CliValFlashSize)
	flashWriteLog.Info().Str("usage", CliValTextFlashOffset).Bool("req", false).Msg(CliValFlashOffset)
	flashWriteLog.Info().Str("usage", CliValTextFlashCompress).Bool("req", false).Msg(CliValFlashCompress)
	fmt.Println()

	//

	obj.End()
}

func (obj *MethodObj) Version() {
	newLogObj := obj.log.Info()

	newLogObj.Str("name", esptool.GlobalName)
	newLogObj.Str("description", "ESP32 flashing utility written in GoLang")
	newLogObj.Str("ver", esptool.GlobalVersion)
	newLogObj.Str("upd", esptool.GlobalDateUpdate)
	newLogObj.Str("hash", esptool.GlobalHash[:8])

	newLogObj.Msg("Version")
	obj.End()
}

func (obj *MethodObj) List() {
	newLog := obj.log.NewLog("List")

	list, err := cmd.ListSerial()
	if err != nil {
		newLog.Debug().Err(err).Send()
		newLog.Warn().Msg("Devices not found")
		obj.End()
	}

	for _, serialDev := range list {
		newLog.Info().Msg(serialDev)
	}

	obj.End()
}
