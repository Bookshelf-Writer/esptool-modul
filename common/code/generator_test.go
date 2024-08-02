//go:build ignore

package code

import (
	"esptool/common/generator"
	"testing"
)

//###########################################################//

func TestErrorCode(t *testing.T) {
	obj := generator.Init("Err", "error.go")
	val := obj.GetByteVal()

	//

	val.Add(0x05, "Received message is invalid")
	val.Add(0x06, "Failed to act on received message")
	val.Add(0x07, "Invalid CRC in message").Delim()
	val.Add(0x08, "Flash write error")
	val.Add(0x09, "Flash read error")
	val.Add(0x0A, "Flash read length error").Delim()
	val.Add(0x0B, "Deflate error")

	//

	build(t, obj)
}

func TestOpCode(t *testing.T) {
	obj := generator.Init("Op", "op.go")
	val := obj.GetByteVal()

	//

	//Commands supported by ESP8266 ROM bootloader
	val.Add(0x02, "Flash Begin")
	val.Add(0x03, "Flash Data")
	val.Add(0x04, "Flash End")
	val.Add(0x05, "Memory Begin")
	val.Add(0x06, "Memory End")
	val.Add(0x07, "Memory Data")
	val.Add(0x08, "Sync")
	val.Add(0x09, "Write Register")
	val.Add(0x0A, "Read Register").Delim()

	//Some commands supported by ESP32 and later chips ROM bootloader (or -8266 w/ stub)
	val.Add(0x0B, "Spi Set Params")
	val.Add(0x0D, "Spi Attach Flash")
	val.Add(0x0E, "Read Flash")
	val.Add(0x0F, "Change Baudrate")
	val.Add(0x10, "Flash Deflate Begin")
	val.Add(0x11, "Flash Deflate lData")
	val.Add(0x12, "Flash Deflate lEnd")
	val.Add(0x13, "Spi Flash MD5").Delim()

	//Some commands supported by stub only
	val.Add(0xD0, "Erase Flash")
	val.Add(0xD1, "Erase Region")
	val.Add(0xD2, "Read Flash Fast")
	val.Add(0xD3, "Run User Code").Delim()

	val.Add(0xD4, "Flash Encrypt Data")

	//

	build(t, obj)
}
