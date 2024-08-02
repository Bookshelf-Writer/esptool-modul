package parser

import (
	"fmt"
)

//###########################################################//

func LoadIndex() (*map[string]string, error) {
	body, err := Load(UrlPyDir + "__init__.py")
	if err != nil {
		return nil, err
	}

	chipDefs := parseGroup(body, "CHIP_DEFS", [2]rune{'{', '}'})
	if len(chipDefs) < 2 {
		return nil, fmt.Errorf("CHIP_DEFS not found.")
	}

	chipDefMap := parseMapX2(chipDefs)

	return &chipDefMap, nil
}
