package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

//###########################################################//

func executeCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("%s: %v", out.String(), err)
	}
	return out.String(), nil
}

func isDeviceAccessible(device string) bool {
	file, err := os.Open(device)
	if err != nil {
		return false
	}
	file.Close()
	return true
}

//////

func listDevices() ([]string, error) {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "linux":
		cmd = "sh"
		args = append(args, "-c", "ls /dev/tty* | grep -E 'ttyUSB.*|ttyACM.*'")
	case "windows":
		cmd = "powershell"
		args = append(args, "Get-ItemProperty HKLM:\\HARDWARE\\DEVICEMAP\\SERIALCOMM | Select-Object -Property * -ExcludeProperty PSPath, PSParentPath, PSChildName, PSDrive, PSProvider | ForEach-Object { $_.PSObject.Properties.Value }")
	case "darwin":
		cmd = "sh"
		args = append(args, "-c", "ls /dev/tty.*")
	default:
		return nil, fmt.Errorf("unsupported platform")
	}

	devices, err := executeCommand(cmd, args...)
	if err != nil {
		return nil, err
	}

	deviceList := strings.Split(strings.TrimSpace(devices), "\n")

	accessibleDevices := []string{}
	for _, device := range deviceList {
		device = strings.TrimSpace(device)
		if device != "" && isDeviceAccessible(device) {
			accessibleDevices = append(accessibleDevices, device)
		}
	}

	return accessibleDevices, nil
}
