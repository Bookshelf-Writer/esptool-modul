package command

import (
	"bytes"
	"github.com/Bookshelf-Writer/esptool-modul/esp32/code"
	"github.com/Bookshelf-Writer/esptool-modul/esp32/portal"
)

//###########################################################//

type CommandObj struct {
	msg portal.MsgObj
}

func newRequest(opcode code.OpType, data []byte) *CommandObj {
	obj := CommandObj{}

	obj.msg.Direction = code.DirectionRequest
	obj.msg.Opcode = opcode
	obj.msg.Data = data
	obj.msg.Length = len(data)
	obj.msg.Checksum = make([]byte, 4)

	return &obj
}

////

func (c *CommandObj) Bytes() []byte {
	var buffer bytes.Buffer

	buffer.WriteByte(byte(c.msg.Direction))
	buffer.WriteByte(byte(c.msg.Opcode))

	buffer.Write(Number.Uint16(uint16(c.msg.Length)))

	buffer.Write(c.msg.Checksum)
	buffer.Write(c.msg.Data)

	return buffer.Bytes()
}

func (c *CommandObj) Checksum(data []byte) *CommandObj {
	state := uint32(0xEF)

	for _, d := range data {
		state ^= uint32(d)
	}

	c.msg.Checksum = Number.Uint32(state)
	return c
}

func (c *CommandObj) Opcode() string {
	return c.msg.Opcode.String()
}

func (c *CommandObj) OpcodeToByte() byte {
	return byte(c.msg.Opcode)
}
