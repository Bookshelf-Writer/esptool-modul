package command

import (
	"bytes"
	"encoding/binary"
)

//###########################################################//

type bufferObj struct {
	buf bytes.Buffer
}

func initBuffer() *bufferObj {
	b := bufferObj{}
	return &b
}

func (b *bufferObj) Bytes() []byte {
	return b.buf.Bytes()
}

func (b *bufferObj) String() string {
	return b.buf.String()
}

//

func (b *bufferObj) Write(data []byte) *bufferObj {
	b.buf.Write(data)
	return b
}

func (b *bufferObj) Uint16(data uint16) *bufferObj {
	var buf []byte

	binary.BigEndian.PutUint16(buf, data)
	b.buf.Write(buf)

	return b
}

func (b *bufferObj) Uint32(data uint32) *bufferObj {
	var buf []byte

	binary.BigEndian.PutUint32(buf, data)
	b.buf.Write(buf)

	return b
}
