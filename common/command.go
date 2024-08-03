package common

import (
	"bytes"
	"encoding/binary"
	"github.com/Bookshelf-Writer/esptool-modul/esp32/code"
)

type Command struct {
	Direction Direction
	Opcode    code.OpType
	Data      []byte
	Length    int
	Checksum  []byte
}

func (c *Command) ToBytes() []byte {
	var buffer bytes.Buffer

	buffer.WriteByte(byte(c.Opcode))
	buffer.WriteByte(byte(c.Direction))

	var size []byte
	binary.BigEndian.PutUint16(size, uint16(c.Length))
	buffer.Write(size)

	buffer.Write(c.Checksum)
	buffer.Write(c.Data)

	return buffer.Bytes()
}

func NewCommand(opcode code.OpType, data []byte) *Command {
	return &Command{
		Direction: DirectionRequest,
		Opcode:    opcode,
		Data:      data,
		Length:    len(data),
		Checksum:  make([]byte, 4),
	}
}

////

func NewReadRegisterCommand(register uint32) *Command {
	return NewCommand(code.OpReadRegister, Uint32ToBytes(register))
}

func NewSyncCommand() *Command {
	payload := []byte{0x07, 0x07, 0x12, 0x20}
	payload = append(payload, bytes.Repeat([]byte{0x55}, 32)...)

	return NewCommand(code.OpSync, payload)
}

func NewAttachSpiFlashCommand() *Command {
	return NewCommand(code.OpSpiAttachFlash, make([]byte, 8))
}

func NewReadFlashCommand(offset uint32, size uint32) *Command {
	payload := Uint32ToBytes(offset)
	payload = append(payload, Uint32ToBytes(size)...)

	return NewCommand(
		code.OpReadFlash,
		payload,
	)
}

func NewChangeBaudrateCommand(newBaudrate uint32, oldBaudrate uint32) *Command {
	payload := Uint32ToBytes(newBaudrate)
	payload = append(payload, Uint32ToBytes(oldBaudrate)...)

	return NewCommand(
		code.OpChangeBaudrate,
		payload,
	)
}

func NewBeginFlashCommand(eraseSize uint32, numBlocks uint32, blockSize uint32, offset uint32) *Command {
	payload := Uint32ToBytes(eraseSize)
	payload = append(payload, Uint32ToBytes(numBlocks)...)
	payload = append(payload, Uint32ToBytes(blockSize)...)
	payload = append(payload, Uint32ToBytes(offset)...)

	return NewCommand(code.OpFlashBegin, payload)
}

func NewBeginFlashDeflCommand(eraseSize uint32, numBlocks uint32, blockSize uint32, offset uint32) *Command {
	payload := Uint32ToBytes(eraseSize)
	payload = append(payload, Uint32ToBytes(numBlocks)...)
	payload = append(payload, Uint32ToBytes(blockSize)...)
	payload = append(payload, Uint32ToBytes(offset)...)

	return NewCommand(code.OpFlashDeflateBegin, payload)
}

func calculateChecksum(data []byte) []byte {
	state := uint32(0xEF)

	for _, d := range data {
		state ^= uint32(d)
	}
	return Uint32ToBytes(state)
}

func NewFlashDataCommand(data []byte, sequence uint32) *Command {
	checksum := calculateChecksum(data)
	payload := Uint32ToBytes(uint32(len(data)))
	payload = append(payload, Uint32ToBytes(sequence)...)
	payload = append(payload, Uint32ToBytes(0)...)
	payload = append(payload, Uint32ToBytes(0)...)
	payload = append(payload, data...)

	cmd := NewCommand(code.OpFlashData, payload)
	cmd.Checksum = checksum

	return cmd
}

func NewFlashDataDeflCommand(data []byte, sequence uint32) *Command {
	checksum := calculateChecksum(data)
	payload := Uint32ToBytes(uint32(len(data)))
	payload = append(payload, Uint32ToBytes(sequence)...)
	payload = append(payload, Uint32ToBytes(0)...)
	payload = append(payload, Uint32ToBytes(0)...)
	payload = append(payload, data...)

	cmd := NewCommand(code.OpFlashDeflateLData, payload)
	cmd.Checksum = checksum

	return cmd
}

func NewFlashEndCommand(reboot bool) *Command {
	param := uint32(0)
	if reboot {
		param = 1
	}
	return NewCommand(
		code.OpFlashEnd,
		Uint32ToBytes(param),
	)
}
