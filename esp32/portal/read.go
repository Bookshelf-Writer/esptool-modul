package portal

import (
	"bytes"
	"github.com/Bookshelf-Writer/esptool-modul/esp32/code"
	"github.com/Bookshelf-Writer/esptool-modul/lib/serial"
	"time"
)

//###########################################################//

func Read(port *serial.SerialObj, timeout time.Duration) ([]byte, error) {
	state := code.StateWaitingHeader
	startTime := time.Now()
	var buf bytes.Buffer

	err := port.Timeout.Set(timeout)
	if err != nil {
		return nil, err
	}

	for {
		if time.Since(startTime) > timeout {
			return nil, ErrTimeoutRead
		}

		byteBuf := make([]byte, 1)
		n, err := port.Read(byteBuf)
		if err != nil {
			if err.Error() == "EOF" {
				continue
			}
			return nil, err
		}
		if n != 1 {
			continue
		}

		switch state {

		case code.StateWaitingHeader:
			if byteBuf[0] == code.SlipHeader.Byte() {
				state = code.StateReadingContent
			}

		case code.StateReadingContent:
			switch byteBuf[0] {
			case code.SlipHeader.Byte():
				return buf.Bytes(), nil

			case code.SlipEscapeChar.Byte():
				state = code.StateInEscape
			default:
				buf.WriteByte(byteBuf[0])
			}

		case code.StateInEscape:
			switch byteBuf[0] {
			case 0xDC:
				buf.WriteByte(code.SlipHeader.Byte())
				state = code.StateReadingContent
			case 0xDD:
				buf.WriteByte(code.SlipEscapeChar.Byte())
				state = code.StateReadingContent
			default:
				return nil, ErrUnexpectedChar
			}

		}
	}
}
