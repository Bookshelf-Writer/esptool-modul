package command

import (
	"github.com/Bookshelf-Writer/esptool-modul/esp32/code"
)

//###########################################################//

type ReadCommandObj struct{}

var Read ReadCommandObj

////

func (ReadCommandObj) Register(register uint32) *CommandObj {
	buf := initBuffer()
	buf.Uint32(register)
	return newRequest(code.OpReadRegister, buf.Bytes())
}

func (ReadCommandObj) Flash(offset uint32, size uint32) *CommandObj {
	buf := initBuffer()

	buf.Uint32(offset)
	buf.Uint32(size)

	return newRequest(code.OpFlashReadSlow, buf.Bytes())
}
