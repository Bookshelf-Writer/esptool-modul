package main

import (
	"esptool/common/generator"
	"esptool/esp32/core"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//###########################################################//

func buildGO(maps map[string]*core.ModulStruct, namespace map[string]string) {
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

		obj.Print("var ").Print(key).PrintLN(" = ModulStruct{")
		obj.Repeat(1).Print("Name: ").PrintString(mod.Name).PrintLN(",").LN()

		obj.Repeat(1).PrintLN("Sys: ModulSystemStruct{")
		obj.Repeat(2).Print("UF2: " + strconv.FormatUint(mod.Sys.UF2, 10)).PrintLN("," + fmt.Sprintf(" //0x%02x", mod.Sys.UF2))
		obj.Repeat(2).Print("Chip: " + strconv.Itoa(mod.Sys.Chip)).PrintLN(",")
		obj.Repeat(2).Print("LenStatus: " + strconv.Itoa(mod.Sys.LenStatus)).PrintLN(",")
		obj.Repeat(2).Print("FlashOffset: " + strconv.FormatUint(mod.Sys.FlashOffset, 10)).PrintLN(",")
		obj.Repeat(1).PrintLN("},")

		obj.Repeat(1).PrintLN("Encrypt: ModulEncryptStruct{")
		obj.Repeat(2).Print("WriteAlign: " + strconv.Itoa(mod.Encrypt.WriteAlign)).PrintLN(",")
		obj.Repeat(2).Print("Supports: " + strconv.FormatBool(mod.Encrypt.Supports)).PrintLN(",")
		obj.Repeat(1).PrintLN("},")

		obj.Repeat(1).PrintLN("MagicValue: []uint64{")
		for _, number := range mod.MagicValue {
			obj.Repeat(2).Print(strconv.FormatUint(number, 10) + ",")
			obj.PrintLN(fmt.Sprintf(" //0x%02x", number))
		}
		obj.Repeat(1).PrintLN("},")

		var bufMapKeys []string
		bufKey := make(map[string]string)
		for keyMemo, objMemo := range mod.Memory {
			kk := fmt.Sprintf("%d:%d:%s", objMemo.Start, objMemo.End, keyMemo)
			bufKey[kk] = keyMemo
			bufMapKeys = append(bufMapKeys, kk)
		}
		sort.Strings(bufMapKeys)

		obj.Repeat(1).PrintLN("Memory: map[string]ModulMapStruct{")
		for _, posKey := range bufMapKeys {
			keyMemo := bufKey[posKey]
			objMemo := mod.Memory[keyMemo]

			obj.Repeat(2).PrintString(keyMemo).Print(": {")
			obj.Print(strconv.FormatUint(objMemo.Start, 10) + ",")
			obj.Print(strconv.FormatUint(objMemo.End, 10))
			obj.Print("},").PrintLN(fmt.Sprintf(" //0x%02x 0x%02x", objMemo.Start, objMemo.End))
		}
		obj.Repeat(1).PrintLN("},")

		obj.PrintLN("}")
		obj.SaveFileBuf("core")
	}

	obj := generator.Init("MAP", "esp32/core/map.go")
	obj.SaveFileBuf("core")

	fmt.Println("GO generated")
}
