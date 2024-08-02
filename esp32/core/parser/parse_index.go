package parser

import (
	"fmt"
	"regexp"
	"strings"
)

//###########################################################//

func LoadIndex() (*map[string]string, error) {
	body, err := Load(UrlPyDir + "__init__.py")
	if err != nil {
		return nil, err
	}

	source := string(body)
	re := regexp.MustCompile(`CHIP_DEFS = \{([^}]+)\}`)
	match := re.FindStringSubmatch(source)

	if len(match) < 2 {
		return nil, fmt.Errorf("CHIP_DEFS not found.")
	}

	chipDefs := strings.ReplaceAll(match[1], "\n", "")
	chipDefs = strings.ReplaceAll(chipDefs, " ", "")

	chipDefMap := make(map[string]string)
	pairs := strings.Split(chipDefs, ",")

	for _, pair := range pairs {
		keyValue := strings.Split(pair, ":")
		if len(keyValue) == 2 {
			key := strings.Trim(keyValue[0], `"`)
			value := keyValue[1]
			chipDefMap[key] = value
		}
	}

	return &chipDefMap, nil
}
