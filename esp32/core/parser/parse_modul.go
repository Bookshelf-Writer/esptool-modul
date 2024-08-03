package main

import (
	"fmt"
	"github.com/Bookshelf-Writer/esptool-modul/esp32/core"
	"regexp"
	"strconv"
	"strings"
)

//###########################################################//

func LoadModul(file string) (*core.ModulStruct, error) {
	body, err := Load(UrlPyDir + file + ".py")
	if err != nil {
		return nil, err
	}

	// Конвертация данных в строку
	content := string(body)

	// Регулярное выражение для нахождения нужного блока текста
	classRegex := regexp.MustCompile(`class \w+\(\w+\):\s*([\s\S]*)`)
	match := classRegex.FindStringSubmatch(content)

	if match == nil || len(match) < 2 {
		return nil, fmt.Errorf("Required pattern not found.")
	}

	classContent := match[1]

	// Поиск всех переменных формата "ключ = значение"
	varRegex := regexp.MustCompile(`([A-Z0-9_]{4,})\s*=\s*([^#\n]*)`)
	varMatches := varRegex.FindAllStringSubmatch(classContent, -1)

	if varMatches == nil {
		return nil, fmt.Errorf("No variables found.")
	}

	obj := core.ModulStruct{}
	bufValueMap := make(map[string]string)
	obj.Value = &bufValueMap

	for _, m := range varMatches {
		if len(m) == 3 {
			key := strings.TrimSpace(m[1])
			val := strings.TrimSpace(m[2])
			if len(val) > 0 {
				switch key {

				case "CHIP_NAME":
					if len(obj.Name) == 0 {
						val = strings.ReplaceAll(val, "\"", "")
						obj.Name = val
					}

				case "CHIP_DETECT_MAGIC_VALUE":
					obj.MagicValue = parseHexValues(val)

				case "STATUS_BYTES_LENGTH":
					obj.Sys.LenStatus, _ = strconv.Atoi(val)
				case "IMAGE_CHIP_ID":
					obj.Sys.Chip, _ = strconv.Atoi(val)
				case "BOOTLOADER_FLASH_OFFSET":
					obj.Sys.FlashOffset = parseUint(val)
				case "UF2_FAMILY_ID":
					obj.Sys.UF2 = parseUint(val)

				case "FLASH_ENCRYPTED_WRITE_ALIGN":
					obj.Encrypt.WriteAlign, _ = strconv.Atoi(val)

				case "SUPPORTS_ENCRYPTED_FLASH":
					obj.Encrypt.Supports, _ = strconv.ParseBool(val)

				default:
					if val[0] != '[' && val[0] != '{' && val[0] != '(' {
						bufValueMap[key] = val
					}
				}
			}
		}
	}

	memoMap := parseGroup(body, "MEMORY_MAP", [2]rune{'[', ']'})
	if len(memoMap) > 0 {
		obj.Memory = parseMemoryMap(memoMap)
	}

	return &obj, nil
}
