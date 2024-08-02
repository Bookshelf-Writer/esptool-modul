package cmd

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

//###########################################################//

func isAccessible(device string) bool {
	file, err := os.Open(device)
	if err != nil {
		return false
	}
	file.Close()
	return true
}

/* getting a list of available SERIAL devices */
func ListSerial() ([]string, error) {
	var cmd string
	var args []string

	switch runtime.GOOS {

	case "linux":
		cmd = "sh"
		args = append(args, "-c", ParamLoadSerialLinux)

	case "windows":
		cmd = "powershell"
		args = append(args, ParamLoadSerialWindows)

	case "darwin":
		cmd = "sh"
		args = append(args, "-c", ParamLoadSerialMac)

	default:
		return nil, fmt.Errorf("unsupported platform")
	}

	devices, err := Run(cmd, args...)
	if err != nil {
		return nil, err
	}

	deviceList := strings.Split(strings.TrimSpace(devices), "\n")
	var accessibleDevices []string

	for _, device := range deviceList {
		device = strings.TrimSpace(device)
		if device != "" && isAccessible(device) {
			accessibleDevices = append(accessibleDevices, device)
		}
	}

	return accessibleDevices, nil
}
