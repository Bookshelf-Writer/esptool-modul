package esptool

import (
	"github.com/Bookshelf-Writer/esptool-modul/common/output"
	"github.com/Bookshelf-Writer/esptool-modul/common/serial"
	"github.com/Bookshelf-Writer/esptool-modul/esp32"
)

func ConnectEsp32(portPath string, connectBaudrate uint32, transferBaudrate uint32, retries uint, logger *output.LogObj) (*esp32.ESP32ROM, error) {
	logger = logger.NewLog("ConnectESP")

	serialPort, err := serial.PortInit(serial.ConfigInit(portPath, connectBaudrate))
	if err != nil {
		logger.Trace().Err(err).Msg("serial.PortInit")
		return nil, err
	}

	esp32 := esp32.NewESP32ROM(serialPort, logger)
	err = esp32.Connect(retries)
	if err != nil {
		logger.Trace().Err(err).Msg("esp32.Connect")
		return nil, err
	}

	return esp32, esp32.ChangeBaudrate(transferBaudrate)
}
