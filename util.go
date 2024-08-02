package main

import (
	"esptool/common/serial"
	"esptool/esp32"
	"fmt"
)

func bold(s string) string {
	return fmt.Sprintf("\033[1m%s\033[0m", s)
}

func underline(s string) string {
	return fmt.Sprintf("\033[4m%s\033[0m", s)
}

func connectEsp32(portPath string, connectBaudrate uint32, transferBaudrate uint32, retries uint, logger *LogObj) (*esp32.ESP32ROM, error) {
	logger = logger.NewLog("ConnectESP")

	serialPort, err := serial.PortInit(serial.ConfigInit(portPath, connectBaudrate))
	if err != nil {
		return nil, fmt.Errorf("Failed to open serial port: %s", err.Error())
	}
	esp32 := esp32.NewESP32ROM(serialPort, &logger.log)
	err = esp32.Connect(retries)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to ESP32: %s", err.Error())
	}
	return esp32, esp32.ChangeBaudrate(transferBaudrate)
}
