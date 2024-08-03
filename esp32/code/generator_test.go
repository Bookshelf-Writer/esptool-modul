package code

import (
	"github.com/Bookshelf-Writer/esptool-modul/common/generator"
	"testing"
)

//###########################################################//

func TestFeatureCode(t *testing.T) {
	obj := generator.Init("Feature", "feature.go")
	val := obj.NewVal()

	//

	val.Add(1, "WiFi")
	val.Add(2, "Bluetooth").Delim()

	val.Add(31, "Single Core")
	val.Add(32, "Dual Core").Delim()

	val.Add(51, "Clock 80MHz")
	val.Add(52, "Clock 160MHz")
	val.Add(53, "Clock 240MHz")
	val.Add(54, "Clock 320MHz").Delim()

	val.Add(101, "Embedded Flash")
	val.Add(102, "VRef calibration")
	val.Add(103, "BLK3 partially reserved").Delim()

	val.Add(151, "Coding Scheme None")
	val.Add(152, "Coding Scheme 3/4")
	val.Add(153, "Coding Scheme Repeat")
	val.Add(154, "Coding Scheme Invalid")

	//

	build(t, obj, val)
}

////////////////

func TestEspCode(t *testing.T) {
	obj := generator.Init("Esp", "esp.go")
	val := obj.NewVal()

	//

	val.Add(0x00, "ESP32 D0-WD Q6")
	val.Add(0x01, "ESP32 D0-WD Q5")
	val.Add(0x02, "ESP32 D2-WD Q5")
	val.Add(0x05, "ESP32 PICO D4")

	//

	build(t, obj, val)
}

////////////////

func TestErrorCode(t *testing.T) {
	obj := generator.Init("Err", "error.go")
	val := obj.NewVal()

	//

	val.Add(0x05, "Received message is invalid")
	val.Add(0x06, "Failed to act on received message")
	val.Add(0x07, "Invalid CRC in message").Delim()
	val.Add(0x08, "Flash write error")
	val.Add(0x09, "Flash read error")
	val.Add(0x0A, "Flash read length error").Delim()
	val.Add(0x0B, "Deflate error")

	//

	build(t, obj, val)
}

////////////////

func TestOpCode(t *testing.T) {
	obj := generator.Init("Op", "op.go")
	val := obj.NewVal()

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

	build(t, obj, val)
}

////////////////

func TestFlashSizeCode(t *testing.T) {
	obj := generator.Init("Size", "size.go")
	val := obj.NewVal()

	//

	val.Add(0x12, "256Kb J1")
	val.Add(0x13, "512Kb J1")
	val.Add(0x14, "1Mb J1")
	val.Add(0x15, "2Mb J1")
	val.Add(0x16, "4Mb J1")
	val.Add(0x17, "8Mb J1")
	val.Add(0x18, "16Mb J1")
	val.Add(0x19, "32Mb J1")
	val.Add(0x1A, "64Mb J1")
	val.Add(0x1B, "128Mb J1")
	val.Add(0x1C, "256Mb J1").Delim()

	val.Add(0x20, "64Mb J2")
	val.Add(0x21, "128Mb J2")
	val.Add(0x22, "256Mb J2").Delim()

	val.Add(0x32, "256Kb J3")
	val.Add(0x33, "512Kb J3")
	val.Add(0x34, "1Mb J3")
	val.Add(0x35, "2Mb J3")
	val.Add(0x36, "4Mb J3")
	val.Add(0x37, "8Mb J3")
	val.Add(0x38, "16Mb J3")
	val.Add(0x39, "32Mb J3")
	val.Add(0x3A, "64Mb J3")

	//

	build(t, obj, val)
}

////////////////

func TestDirectionCode(t *testing.T) {
	obj := generator.Init("Direction", "direction.go")
	val := obj.NewVal()

	//

	val.Add(0x00, "Request")
	val.Add(0x01, "Response")

	//

	build(t, obj, val)
}
