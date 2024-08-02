package main

import (
	"esptool/esp32/core"
	"fmt"
	"os"
	"strings"
)

//###########################################################//

func main() {
	maps, err := LoadIndex()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	list := make(map[string]*core.ModulStruct)
	for key, value := range *maps {
		if strings.Contains(key, "beta") {
			continue
		}

		fmt.Println(key, value)
		obj, err := LoadModul(key)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		obj.Value = nil //clear other variables
		list[value] = obj
	}

	fmt.Println("")
	buildJSON(list)
	buildGO(list, *maps)
}
