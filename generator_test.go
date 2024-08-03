package main

import (
	"esptool/common/generator"
	"testing"
)

//###########################################################//

func TestCliTrig(t *testing.T) {
	obj := generator.Init("CliTrig", "cli_trig.go")
	val := obj.NewVal()

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

	obj.PrintLN("import \"flag\"").LN()
	build(t, obj)

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
		obj.Repeat(1).TitleCase(code).Print(": flag.Bool( ")
		obj.ConstCode(code).Print(", false, ")
		obj.ConstText(code).PrintLN("),")
	}
	obj.PrintLN("}").LN()

	//

	obj.Print("var ").Type().PrintLN("Map = map[string]*bool{")
	for _, code := range obj.GetStrings() {
		obj.Repeat(1).ConstCode(code).Print(": ")
		obj.Type().Print(".").TitleCase(code).PrintLN(",")
	}
	obj.PrintLN("}")

	//

	err := obj.SaveFileBuf("main")
	if err != nil {
		t.Fatal(err)
	}

}

func TestCliValue(t *testing.T) {
	obj := generator.Init("CliVal", "cli_val.go")
	val := obj.NewVal()

	//

	val.Add("port", "Device name or path")

	val.Add("baudConnect", "Serial signalling rate during connect phase")
	val.Add("baudTransfer", "Serial signalling rate during data transfer").Delim()

	val.Add("connTimeout", "Timeout to wait for chip response upon connecting")
	val.Add("connRetries", "How often to retry connecting").Delim()

	val.Add("flashCompress", "Use compression for transfer")
	val.Add("flashOffset", "The point where we start")
	val.Add("flashSize", "How many bytes do we capture?")
	val.Add("flashFile", "File path")

	//

	build(t, obj)

	//

	err := obj.SaveFileBuf("main")
	if err != nil {
		t.Fatal(err)
	}

}
