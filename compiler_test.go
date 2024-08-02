package main

import (
	"esptool/common/generator"
	"testing"
)

func build(t *testing.T, obj *generator.GeneratorObj) {

	obj.PrintLN("const (")
	for _, code := range obj.GetStrings() {
		obj.Repeat(1).ConstCode(code).Print(" ").Print(" = ").PrintString(code).LN()
		obj.SetDelimString(code)
	}
	obj.PrintLN(")").LN()

	//

	obj.PrintLN("const (")
	for _, code := range obj.GetStrings() {
		obj.Repeat(1).ConstText(code).Print(" = ").PrintString(obj.GetTextString(code)).LN()
		obj.SetDelimString(code)
	}
	obj.PrintLN(")").LN()

}
