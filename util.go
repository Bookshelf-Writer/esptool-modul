package esptool

import (
	"github.com/Bookshelf-Writer/esptool-modul/common/output"
	"github.com/Bookshelf-Writer/esptool-modul/common/serial"
	"github.com/Bookshelf-Writer/esptool-modul/esp32"
	"github.com/Bookshelf-Writer/esptool-modul/esp32/command"
	"time"
)

func ConnectEsp32(portPath string, connectBaudrate uint32, transferBaudrate uint32, retries uint, logger *output.LogObj) (*esp32.ESP32ROM, error) {
	logger = logger.NewLog("ConnectESP")

	serialPort, err := serial.PortInit(serial.ConfigInit(portPath, connectBaudrate))
	if err != nil {
		logger.Trace().Err(err).Msg("serial.PortInit")
		return nil, err
	}

	esp := esp32.NewESP32ROM(serialPort, logger)
	err = esp.Connect(retries)
	if err != nil {
		logger.Trace().Err(err).Msg("esp32.Connect")
		return nil, err
	}

	//установка скорости подключения
	{
		_, err = esp32.CheckExecuteCommand(serialPort,
			command.ChangeBaudRate(transferBaudrate, 0),
			100*time.Millisecond,
			3,
		)
		if err != nil {
			return nil, err
		}

		err = serialPort.BaudRate.Set(transferBaudrate)
		if err != nil {
			return nil, err
		}

		time.Sleep(10 * time.Millisecond)

		err = serialPort.Flush() // get rid of crap sent during baud rate change
	}

	return esp, err
}
