package command

import (
	"bytes"
	"github.com/Bookshelf-Writer/esptool-modul/esp32/code"
)

//###########################################################//

type CommandObj struct {
	direction code.DirectionType
	opcode    code.OpType
	data      []byte
	length    int
	checksum  []byte
}

func newRequest(opcode code.OpType, data []byte) *CommandObj {
	return &CommandObj{
		direction: code.DirectionRequest,
		opcode:    opcode,
		data:      data,
		length:    len(data),
		checksum:  make([]byte, 4),
	}
}

////

func (c *CommandObj) Bytes() []byte {
	var buffer bytes.Buffer

	buffer.WriteByte(byte(c.direction))
	buffer.WriteByte(byte(c.opcode))

	buffer.Write(initBuffer().Uint16(uint16(c.length)).Bytes())

	buffer.Write(c.checksum)
	buffer.Write(c.data)

	return buffer.Bytes()
}

func (c *CommandObj) Checksum(data []byte) *CommandObj {
	state := uint32(0xEF)

	for _, d := range data {
		state ^= uint32(d)
	}

	c.checksum = initBuffer().Uint32(state).Bytes()
	return c
}

func (c *CommandObj) Opcode() string {
	return c.opcode.String()
}

func (c *CommandObj) OpcodeToByte() byte {
	return byte(c.opcode)
}
