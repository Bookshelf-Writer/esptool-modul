package code

import (
	"testing"
)

//###########################################################//

func TestErrorCode(t *testing.T) {
	obj := generatorInit("Err", "error.go")

	//

	obj.val.Add(0x05, "Received message is invalid")
	obj.val.Add(0x06, "Failed to act on received message")
	obj.val.Add(0x07, "Invalid CRC in message").Delim()
	obj.val.Add(0x08, "Flash write error")
	obj.val.Add(0x09, "Flash read error")
	obj.val.Add(0x0A, "Flash read length error").Delim()
	obj.val.Add(0x0B, "Deflate error")

	//

	obj.Build(t)
}

func TestOpCode(t *testing.T) {
	obj := generatorInit("Op", "op.go")

	//

	obj.val.Add(0x02, "Flash Begin")
	obj.val.Add(0x03, "Flash Data")
	obj.val.Add(0x04, "Flash End")
	obj.val.Add(0x05, "Memory Begin")
	obj.val.Add(0x06, "Memory End")
	obj.val.Add(0x07, "Memory Data")
	obj.val.Add(0x08, "Sync")
	obj.val.Add(0x09, "Write Register")
	obj.val.Add(0x0A, "Read Register").Delim()

	obj.val.Add(0x0B, "Spi Set Params")
	obj.val.Add(0x0D, "Spi Attach Flash")
	obj.val.Add(0x0E, "Read Flash")
	obj.val.Add(0x0F, "Change Baudrate")
	obj.val.Add(0x10, "Flash Deflate Begin")
	obj.val.Add(0x11, "Flash Deflate lData")
	obj.val.Add(0x12, "Flash Deflate lEnd")
	obj.val.Add(0x13, "Spi Flash MD5").Delim()

	obj.val.Add(0xD0, "Erase Flash")
	obj.val.Add(0xD1, "Erase Region")
	obj.val.Add(0xD2, "Read Flash Fast")
	obj.val.Add(0xD3, "Run User Code").Delim()

	//

	obj.Build(t)
}
