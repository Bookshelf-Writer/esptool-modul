package main

import (
	generator2 "github.com/Bookshelf-Writer/esptool-modul/lib/generator"
	"testing"
)

func build(t *testing.T, obj *generator2.GeneratorObj, val *generator2.GeneratorValueObj) {

	obj.PrintLN("const (")
	for _, code := range val.Get.Strings() {
		obj.Offset(1).Name.SelfCode(code).Print(" ").Print(" = ").String(code).LN()

		if val.Get.IsDelim(code) {
			obj.LN()
		}
	}
	obj.PrintLN(")").LN()

	//

	obj.PrintLN("const (")
	for _, code := range val.Get.Strings() {
		obj.Offset(1).Name.TextCode(code).Print(" = ").String(val.Get.Text(code)).LN()

		if val.Get.IsDelim(code) {
			obj.LN()
		}
	}
	obj.PrintLN(")").LN()

}
