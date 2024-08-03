package main

import (
	"github.com/Bookshelf-Writer/esptool-modul/common/generator"
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
	val.Add("noColor", "Display output without colorization").Delim()

	val.Add("list", "Get a list of available COM devices")
	val.Add("info", "Show ESP information")
	val.Add("flashRead", "Reading firmware from ESP to file")
	val.Add("flashWrite", "Recording firmware from a file in ESP")

	//

	build(t, obj, val)

	//

	obj.Print("type ").Name.Type().PrintLN("Obj struct {")
	for _, code := range val.Get.Strings() {
		obj.Offset(1).PrintLN(obj.Name.ToTitleCase(code) + " *bool")

		if val.Get.IsDelim(code) {
			obj.LN()
		}
	}
	obj.PrintLN("}").LN()

	//

	obj.Print("var ").Name.Type().Print(" = ").Name.Type().PrintLN("Obj{")
	for _, code := range val.Get.Strings() {
		obj.Offset(1).Print(obj.Name.ToTitleCase(code) + ": flag.Bool( ")
		obj.Name.SelfCode(code).Print(", false, ")
		obj.Name.TextCode(code).PrintLN("),")
	}
	obj.PrintLN("}").LN()

	//

	obj.Print("var ").Name.Type().PrintLN("Map = map[string]*bool{")
	for _, code := range val.Get.Strings() {
		obj.Offset(1).Name.SelfCode(code).Print(": ")
		obj.Name.Type().PrintLN("." + obj.Name.ToTitleCase(code) + ",")
	}
	obj.PrintLN("}")

	//

	err := obj.Save("main").Add.Import("flag").End()
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

	build(t, obj, val)

	//

	err := obj.Save("main").End()
	if err != nil {
		t.Fatal(err)
	}

}
