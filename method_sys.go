package main

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

	obj.End()
}
