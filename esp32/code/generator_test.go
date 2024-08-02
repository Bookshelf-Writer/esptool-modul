package code

import (
	"esptool/common/generator"
	"testing"
)

//###########################################################//

func TestFeatureCode(t *testing.T) {
	obj := generator.Init("Feature", "feature.go")
	val := obj.GetVal()

	//

	val.Add(1, "WiFi")
	val.Add(2, "Bluetooth").Delim()

	val.Add(31, "Single Core")
	val.Add(32, "Dual Core").Delim()

	val.Add(51, "80MHz")
	val.Add(52, "160MHz")
	val.Add(53, "240MHz")
	val.Add(54, "320MHz").Delim()

	val.Add(101, "Embedded Flash")
	val.Add(102, "VRef calibration")
	val.Add(103, "BLK3 partially reserved").Delim()

	val.Add(151, "Coding Scheme None")
	val.Add(152, "Coding Scheme 3/4")
	val.Add(153, "Coding Scheme Repeat")
	val.Add(154, "Coding Scheme Invalid")

	//

	build(t, obj)
}
