package portal

import "github.com/Bookshelf-Writer/esptool-modul/esp32/code"

//###########################################################//

type MsgObj struct {
	Direction code.DirectionType
	Opcode    code.OpType

	Data   []byte
	Length int

	Checksum []byte
}
