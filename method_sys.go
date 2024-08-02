package main

import "esptool/common/cmd"

//###########################################################//

func (obj *MethodObj) Help() {

	obj.End()
}

func (obj *MethodObj) Version() {
	newLogObj := obj.log.Info()

	newLogObj.Str("name", GlobalName)
	newLogObj.Str("description", "ESP32 flashing utility written in GoLang")
	newLogObj.Str("ver", GlobalVersion)
	newLogObj.Str("upd", GlobalDateUpdate)
	newLogObj.Str("hash", GlobalHash[:8])

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
