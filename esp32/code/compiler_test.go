package code

import (
	generator2 "github.com/Bookshelf-Writer/esptool-modul/lib/generator"
	"testing"
)

func build(t *testing.T, obj *generator2.GeneratorObj, val *generator2.GeneratorValueObj) {

	obj.PrintLN("const (")
	for _, code := range val.Get.Ints() {
		text := val.Get.Text(code)

		obj.Offset(1).Name.SelfCode(text).Print(" ")
		obj.Name.Type().Print(" = ")
		obj.Number(code).Print("\t//").Hex(code).LN()

		if val.Get.IsDelim(code) {
			obj.LN()
		}
	}
	obj.PrintLN(")").LN()

	//

	obj.PrintLN("const (")
	for _, code := range val.Get.Ints() {
		text := val.Get.Text(code)

		obj.Offset(1).Name.TextCode(text).Print(" = ")
		obj.String(text).LN()

		if val.Get.IsDelim(code) {
			obj.LN()
		}
	}
	obj.PrintLN(")").LN()

	//

	obj.Print("var ").Name.Map().Print(" = map[").Name.Type().PrintLN("]string{")
	for _, code := range val.Get.Ints() {
		text := val.Get.Text(code)

		obj.Offset(1).Name.SelfCode(text).Print(": ")
		obj.Name.TextCode(text).PrintLN(",")

	}
	obj.PrintLN("}").LN()

	obj.Print("func (obj ").Name.Type().PrintLN(") String() string {")
	obj.Offset(1).Print("val, ok := ").Name.Map().PrintLN("[obj]")
	obj.Offset(1).PrintLN("if ok {").Offset(2).PrintLN("return val").Offset(1).PrintLN("}")
	obj.Offset(1).Print("return \"Unknown ").Name.Type().PrintLN("\"")
	obj.PrintLN("}")

	//

	err := obj.Save("code").Add.Type(obj.Name.GetType(), "byte").End()
	if err != nil {
		t.Fatal(err)
	}
}
