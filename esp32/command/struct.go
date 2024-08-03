package command

import (
	"bytes"
	"encoding/binary"
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

	buffer.WriteByte(byte(c.opcode))
	buffer.WriteByte(byte(c.direction))

	var size []byte
	binary.BigEndian.PutUint16(size, uint16(c.length))
	buffer.Write(size)

	buffer.Write(c.checksum)
	buffer.Write(c.data)

	return buffer.Bytes()
}

func (c *CommandObj) Checksum(data []byte) *CommandObj {
	state := uint32(0xEF)

	for _, d := range data {
		state ^= uint32(d)
	}

	binary.BigEndian.PutUint32(c.checksum, state)
	return c
}

func (c *CommandObj) Opcode() string {
	return c.opcode.String()
}
