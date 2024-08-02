package cmd

import (
	"esptool/common/serial"
	"fmt"
	"runtime"
	"strings"
)

//###########################################################//

/*Checking device availability */
func IsAccessible(device string) bool {
	serialPort, err := serial.PortInit(serial.ConfigInit(device, 115200))
	if err != nil {
		return false
	}

	serialPort.Close()
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
		if device != "" && IsAccessible(device) {
			accessibleDevices = append(accessibleDevices, device)
		}
	}

	return accessibleDevices, nil
}
