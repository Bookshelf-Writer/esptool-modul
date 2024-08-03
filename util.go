package esptool

import (
	"fmt"
	"github.com/Bookshelf-Writer/esptool-modul/common/output"
	"github.com/Bookshelf-Writer/esptool-modul/common/serial"
	"github.com/Bookshelf-Writer/esptool-modul/esp32"
)

func ConnectEsp32(portPath string, connectBaudrate uint32, transferBaudrate uint32, retries uint, logger *output.LogObj) (*esp32.ESP32ROM, error) {
	logger = logger.NewLog("ConnectESP")

	serialPort, err := serial.PortInit(serial.ConfigInit(portPath, connectBaudrate))
	if err != nil {
		logger.Debug().Err(err).Msg("serial.PortInit")
		return nil, fmt.Errorf("Failed to open serial port: %s", err.Error())
	}
	esp32 := esp32.NewESP32ROM(serialPort, logger.ZeroLog())
	err = esp32.Connect(retries)
	if err != nil {
		logger.Debug().Err(err).Msg("esp32.Connect")
		return nil, fmt.Errorf("Failed to connect to ESP32: %s", err.Error())
	}
	return esp32, esp32.ChangeBaudrate(transferBaudrate)
}
