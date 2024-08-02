package code

import (
	"esptool/common/generator"
	"testing"
)

//###########################################################//

func TestErrorCode(t *testing.T) {
	obj := generator.Init("Err", "error.go")
	val := obj.GetVal()

	//

	val.Add(0000, "Received")

	//

	build(t, obj)
}
