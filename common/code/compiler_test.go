//go:build ignore

package code

import (
	"esptool/common/generator"
	"testing"
)

func build(t *testing.T, obj *generator.GeneratorObj) {

	obj.PrintLN("const (")
	for _, code := range obj.GetList() {
		text := obj.GetText(code)

		obj.Repeat(1).ConstCode(text).Print(" ").Type().Print(" = ").Byte(code).LN()

		obj.SetDelim(code)
	}
	obj.PrintLN(")").LN()

	//

	obj.PrintLN("const (")
	for _, code := range obj.GetList() {
		text := obj.GetText(code)

		obj.Repeat(1).ConstText(text).Print(" = ").PrintString(text).LN()

		obj.SetDelim(code)
	}
	obj.PrintLN(")").LN()

	//

	obj.Print("var ").Map().Print(" = map[").Type().PrintLN("]string{")
	for _, code := range obj.GetList() {
		text := obj.GetText(code)

		obj.Repeat(1).ConstCode(text).Print(": ").ConstText(text).PrintLN(",")
	}
	obj.PrintLN("}").LN()

	obj.Print("func (obj ").Type().PrintLN(") String() string {")
	obj.Repeat(1).Print("val, ok := ").Map().PrintLN("[obj]")
	obj.Repeat(1).PrintLN("if ok {").Repeat(2).PrintLN("return val").Repeat(1).PrintLN("}")
	obj.Repeat(1).Print("return \"Unknown ").Type().PrintLN("\"")
	obj.PrintLN("}")

	//

	err := obj.SaveFile("code")
	if err != nil {
		t.Fatal(err)
	}
}
