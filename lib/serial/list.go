package serial

import (
	"bytes"
	"os/exec"
	"runtime"
	"strings"
)

//###########################################################//

func ListSerial() ([]string, error) {
	paramLoadSerial, ok := ParamLoadSerialMap[runtime.GOOS]
	if !ok {
		return nil, ErrListSerialUnsupportedPlatform
	}

	var deviceList []string
	var accessibleDevices []string

	{
		var out bytes.Buffer
		cmd := exec.Command(paramLoadSerial.CMD, paramLoadSerial.Args...)
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			return nil, err
		}

		deviceList = strings.Split(strings.TrimSpace(out.String()), "\n")
	}

	for _, device := range deviceList {
		device = strings.TrimSpace(device)

		if len(device) < 2 {
			continue
		}
		if !Check(device) {
			continue
		}

		accessibleDevices = append(accessibleDevices, device)
	}

	return accessibleDevices, nil
}
