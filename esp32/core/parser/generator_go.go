package main

import (
	"esptool/common/generator"
	"esptool/esp32/core"
	"fmt"
)

//###########################################################//

func buildGO(maps map[string]*core.ModulStruct, namespace map[string]string) {
	for filename, key := range namespace {
		obj := generator.Init(key, "esp32/core/"+filename+".go")

		obj.Print("var ").Print(key).PrintLN(" ModulStruct")

		obj.SaveFileBuf("core")
	}
	fmt.Println("GO generated")
}
