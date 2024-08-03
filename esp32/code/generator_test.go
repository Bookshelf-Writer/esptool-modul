package code

import (
	"esptool/common/generator"
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
