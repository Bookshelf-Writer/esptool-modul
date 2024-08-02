//go:build ignore

package main

import (
	"encoding/json"
	"esptool/esp32/core"
	"fmt"
	"os"
)

//###########################################################//

func buildJSON(maps map[string]*core.ModulStruct) {
	jsonData, err := json.MarshalIndent(maps, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling data:", err)
		return
	}

	err = os.WriteFile("esp32/core/esptool-list.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("JSON generated")
}
