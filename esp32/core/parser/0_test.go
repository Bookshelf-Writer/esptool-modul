package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	maps, err := LoadIndex()
	if err != nil {
		t.Fatal(err)
	}

	list := make(map[string]*ModulStruct)
	for key, value := range *maps {
		if strings.Contains(key, "beta") {
			continue
		}

		fmt.Println(key, value)
		obj, err := LoadModul(key)
		if err != nil {
			t.Fatal(err)
		}

		obj.Value = nil //очистка остальных переменных
		list[value] = obj
	}

	jsonData, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling data:", err)
		return
	}

	// Запис JSON у файл
	err = os.WriteFile("output.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}
