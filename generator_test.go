package main

import (
	"esptool/common/generator"
	"testing"
)

//###########################################################//

func TestCLI(t *testing.T) {
	obj := generator.Init("CliTrig", "cli_trig.go")
	val := obj.GetStringVal()

	//

	val.Add("help", "Show a help page")
	val.Add("version", "Show application build information").Delim()

	val.Add("json", "Output responses in json")
	val.Add("noColor", "Display answers without colorization").Delim()

	val.Add("list", "Get a list of available COM devices")
	val.Add("info", "Show ESP information")
	val.Add("flashRead", "Reading firmware from ESP to file")
	val.Add("flashWrite", "Recording firmware from a file in ESP")

	//

	obj.PrintLN("const (")
	for _, code := range obj.GetStrings() {
		text := obj.GetTextString(code)

		obj.Repeat(1).ConstCode(text).Print(" ").Print(" = ").PrintString(code).LN()

		obj.SetDelimString(code)
	}
	obj.PrintLN(")").LN()

	//

	obj.PrintLN("const (")
	for _, code := range obj.GetStrings() {
		text := obj.GetTextString(code)

		obj.Repeat(1).ConstText(text).Print(" = ").PrintString(text).LN()

		obj.SetDelimString(code)
	}
	obj.PrintLN(")").LN()

	//

	obj.Print("type ").Type().PrintLN("Obj struct {")
	for _, code := range obj.GetStrings() {
		obj.Repeat(1).TitleCase(code).PrintLN(" *bool")
		obj.SetDelimString(code)
	}
	obj.PrintLN("}").LN()

	//

	obj.Print("var ").Type().Print(" = ").Type().PrintLN("Obj{")
	for _, code := range obj.GetStrings() {
		text := obj.GetTextString(code)
		obj.Repeat(1).TitleCase(code).Print(": flag.Bool( ")
		obj.ConstCode(text).Print(", false, ")
		obj.ConstText(text).PrintLN("),")
	}
	obj.PrintLN("}")

	//

	err := obj.SaveFileBuf("main")
	if err != nil {
		t.Fatal(err)
	}

}
