package code

import "testing"

func (obj *generatorObj) Build(t *testing.T) {

	obj.PrintLN("const (")
	for _, code := range obj.val.list {
		text := obj.val.maps[code]

		obj.Repeat(1).ConstCode(text).Print(" ").Type().Print(" = ").Byte(code).LN()

		_, ok := obj.val.delim[code]
		if ok {
			obj.LN()
		}
	}
	obj.PrintLN(")").LN()

	//

	obj.PrintLN("const (")
	for _, code := range obj.val.list {
		text := obj.val.maps[code]

		obj.Repeat(1).ConstText(text).Print(" = ").PrintString(text).LN()

		_, ok := obj.val.delim[code]
		if ok {
			obj.LN()
		}
	}
	obj.PrintLN(")").LN()

	//

	obj.Print("var ").Map().Print(" = map[").Type().PrintLN("]string{")
	for _, code := range obj.val.list {
		text := obj.val.maps[code]

		obj.Repeat(1).ConstCode(text).Print(": ").ConstText(text).PrintLN(",")
	}
	obj.PrintLN("}").LN()

	obj.Print("func (obj ").Type().PrintLN(") String() string {")
	obj.Repeat(1).Print("val, ok := ").Map().PrintLN("[obj]")
	obj.Repeat(1).PrintLN("if ok {").Repeat(2).PrintLN("return val").Repeat(1).PrintLN("}")
	obj.Repeat(1).Print("return \"Unknown ").Type().PrintLN("\"")
	obj.PrintLN("}")

	//

	err := obj.SaveFile()
	if err != nil {
		t.Fatal(err)
	}
}
