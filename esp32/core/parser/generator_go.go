package main

import (
	"fmt"
	"github.com/Bookshelf-Writer/esptool-modul/esp32/core"
	"github.com/Bookshelf-Writer/esptool-modul/lib/generator"
	"sort"
	"strconv"
	"strings"
)

//###########################################################//

func buildGO(maps map[string]*core.ModulStruct, namespace map[string]string) {
	objMap := generator.Init("Esp", "esp32/core/map.go")
	objMap.PrintLN("var EspList = []*ModulStruct{")

	//todo переписать на новый генератор и спользованием val
	for filename, key := range namespace {
		if strings.Contains(key, "beta") {
			continue
		}

		obj := generator.Init(key, "esp32/core/"+filename+".go")
		buf, ok := maps[key]
		if !ok {
			continue
		}
		mod := *buf
		objMap.Offset(1).PrintLN("&" + key + ",")

		obj.Print("var ").Print(key).PrintLN(" = ModulStruct{")
		obj.Offset(1).Print("Name: ").String(mod.Name).PrintLN(",").LN()

		fmt.Println(mod.Name, (byte(mod.Sys.Chip)>>1)&0x07)

		obj.Offset(1).PrintLN("Sys: ModulSystemStruct{")
		obj.Offset(2).Print("UF2: " + strconv.FormatUint(mod.Sys.UF2, 10)).PrintLN("," + fmt.Sprintf(" //0x%02x", mod.Sys.UF2))
		obj.Offset(2).Print("Chip: " + strconv.Itoa(mod.Sys.Chip)).PrintLN(",")
		obj.Offset(2).Print("LenStatus: " + strconv.Itoa(mod.Sys.LenStatus)).PrintLN(",")
		obj.Offset(2).Print("FlashOffset: " + strconv.FormatUint(mod.Sys.FlashOffset, 10)).PrintLN(",")
		obj.Offset(1).PrintLN("},")

		obj.Offset(1).PrintLN("Encrypt: ModulEncryptStruct{")
		obj.Offset(2).Print("WriteAlign: " + strconv.Itoa(mod.Encrypt.WriteAlign)).PrintLN(",")
		obj.Offset(2).Print("Supports: " + strconv.FormatBool(mod.Encrypt.Supports)).PrintLN(",")
		obj.Offset(1).PrintLN("},")

		obj.Offset(1).PrintLN("MagicValue: []uint64{")
		for _, number := range mod.MagicValue {
			obj.Offset(2).Print(strconv.FormatUint(number, 10) + ",")
			obj.PrintLN(fmt.Sprintf(" //0x%02x", number))
		}
		obj.Offset(1).PrintLN("},")

		var bufMapKeys []string
		bufKey := make(map[string]string)
		for keyMemo, objMemo := range mod.Memory {
			kk := fmt.Sprintf("%d:%d:%s", objMemo.Start, objMemo.End, keyMemo)
			bufKey[kk] = keyMemo
			bufMapKeys = append(bufMapKeys, kk)
		}
		sort.Strings(bufMapKeys)

		obj.Offset(1).PrintLN("Memory: map[string]ModulMapStruct{")
		for _, posKey := range bufMapKeys {
			keyMemo := bufKey[posKey]
			objMemo := mod.Memory[keyMemo]

			obj.Offset(2).String(keyMemo).Print(": {")
			obj.Print(strconv.FormatUint(objMemo.Start, 10) + ",")
			obj.Print(strconv.FormatUint(objMemo.End, 10))
			obj.Print("},").PrintLN(fmt.Sprintf(" //0x%02x 0x%02x", objMemo.Start, objMemo.End))
		}
		obj.Offset(1).PrintLN("},")

		obj.PrintLN("}")
		obj.Save("core").End()
	}

	objMap.PrintLN("}")
	objMap.Save("core").End()

	fmt.Println("GO generated")
}
