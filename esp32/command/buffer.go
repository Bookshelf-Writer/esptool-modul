package command

import (
	"bytes"
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
	buf := []byte{byte(data & 0xFF),
		byte((data >> 8) & 0xFF)}

	b.buf.Write(buf)
	return b
}

func (b *bufferObj) Uint32(data uint32) *bufferObj {
	buf := []byte{byte(data & 0xFF),
		byte((data >> 8) & 0xFF),
		byte((data >> 16) & 0xFF),
		byte((data >> 24) & 0xFF),
	}

	b.buf.Write(buf)
	return b
}
