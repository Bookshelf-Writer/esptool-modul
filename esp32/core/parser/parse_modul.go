package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//###########################################################//

type ModulMapStruct struct {
	Start uint64 `json:"start"`
	End   uint64 `json:"end"`
}

type ModulStruct struct {
	Value     *map[string]string `json:"value,omitempty"`
	Name      string             `json:"name"`
	ChipId    int                `json:"chipId"`
	LenStatus int                `json:"lenStatus"`

	FlashEncryptedWriteAlign int `json:"flashEncryptedWriteAlign"`

	IROM ModulMapStruct `json:"IROM"`
	DROM ModulMapStruct `json:"DROM"`

	MagicValue []uint64 `json:"magicValue"`

	BootloaderFlashOffset uint64 `json:"bootloaderFlashOffset"`
	Uf2FamilyId           uint64 `json:"uf2FamilyId"`

	SupportsEncryptedFlash bool `json:"supportsEncryptedFlash"`
	IsStub                 bool `json:"isStub"`

	Memory map[string]ModulMapStruct `json:"memory"`
}

func LoadModul(file string) (*ModulStruct, error) {
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

	obj := ModulStruct{}
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
					obj.LenStatus, _ = strconv.Atoi(val)

				case "FLASH_ENCRYPTED_WRITE_ALIGN":
					obj.FlashEncryptedWriteAlign, _ = strconv.Atoi(val)

				case "IS_STUB":
					obj.IsStub, _ = strconv.ParseBool(val)
				case "SUPPORTS_ENCRYPTED_FLASH":
					obj.SupportsEncryptedFlash, _ = strconv.ParseBool(val)

				case "BOOTLOADER_FLASH_OFFSET":
					obj.BootloaderFlashOffset = parseUint(val)

				case "UF2_FAMILY_ID":
					obj.Uf2FamilyId = parseUint(val)

				case "IROM_MAP_START":
					obj.IROM.Start = parseUint(val)
				case "IROM_MAP_END":
					obj.IROM.End = parseUint(val)

				case "DROM_MAP_START":
					obj.DROM.Start = parseUint(val)
				case "DROM_MAP_END":
					obj.DROM.End = parseUint(val)

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
