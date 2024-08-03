package main

import (
	"fmt"
	"github.com/Bookshelf-Writer/esptool-modul/esp32/core"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

//###########################################################//

func Load(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Received non-200 response code: %d\n", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func parseGroup(data []byte, value string, delim [2]rune) string {
	pat := value + " = \\" + string(delim[0]) + "([^}]+)\\" + string(delim[1])
	re := regexp.MustCompile(pat)
	match := re.FindStringSubmatch(string(data))

	if len(match) < 2 {
		return ""
	}

	chipDefs := strings.ReplaceAll(match[1], "\n", "")
	return strings.ReplaceAll(chipDefs, " ", "")
}

func parseMapX2(text string) map[string]string {
	chipDefMap := make(map[string]string)
	pairs := strings.Split(text, ",")

	for _, pair := range pairs {
		keyValue := strings.Split(pair, ":")
		if len(keyValue) == 2 {
			key := strings.Trim(keyValue[0], `"`)
			value := keyValue[1]
			chipDefMap[key] = value
		}
	}

	return chipDefMap
}

func parseMemoryMap(input string) map[string]core.ModulMapStruct {
	pattern := `\[\s*(0x[0-9A-F]+)\s*,\s*(0x[0-9A-F]+)\s*,\s*"([^"]+)"\s*\]`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(input, -1)

	memoryMap := make(map[string]core.ModulMapStruct)
	for _, match := range matches {
		start, err := strconv.ParseUint(match[1], 0, 64)
		if err != nil {
			return memoryMap
		}
		end, err := strconv.ParseUint(match[2], 0, 64)
		if err != nil {
			return memoryMap
		}
		memoryMap[match[3]] = core.ModulMapStruct{start, end}
	}

	return memoryMap
}

func parseHexValues(input string) []uint64 {
	re := regexp.MustCompile(`0x[0-9A-Fa-f]+`)
	hexMatches := re.FindAllString(input, -1)

	var values []uint64
	for _, hex := range hexMatches {
		value, err := strconv.ParseUint(hex, 0, 64)
		if err != nil {
			return nil
		}
		values = append(values, value)
	}

	return values
}

func parseUint(str string) uint64 {
	val, err := strconv.ParseUint(str, 0, 64)
	if err != nil {
		return 0
	}
	return val
}
