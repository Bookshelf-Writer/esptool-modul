package portal

import (
	"bytes"
	"github.com/Bookshelf-Writer/esptool-modul/esp32/code"
)

//###########################################################//

func encode(data []byte) []byte {
	var buf bytes.Buffer
	buf.WriteByte(code.SlipHeader.Byte())

	for _, b := range data {
		switch b {
		case code.SlipEscapeChar.Byte():
			buf.Write([]byte{code.SlipEscapeChar.Byte(), 0xDD})
		case code.SlipHeader.Byte():
			buf.Write([]byte{code.SlipEscapeChar.Byte(), 0xDC})
		default:
			buf.WriteByte(b)
		}

	}

	buf.WriteByte(code.SlipHeader.Byte())
	return buf.Bytes()
}
