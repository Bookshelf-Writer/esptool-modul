package serial

import (
	"go.bug.st/serial"
)

//###########################################################//

type ConfigObj struct {
	PortName string
	BaudRate uint32
	DataBits int
	StopBits serial.StopBits
	Parity   serial.Parity
}

func ConfigInit(portName string, baudRate uint32) *ConfigObj {
	return &ConfigObj{
		PortName: portName,
		BaudRate: baudRate,
		DataBits: DataBits,
		StopBits: serial.OneStopBit,
		Parity:   serial.NoParity,
	}
}

func (config *ConfigObj) Mode() *serial.Mode {
	return &serial.Mode{
		BaudRate: int(config.BaudRate),
		DataBits: config.DataBits,
		StopBits: config.StopBits,
		Parity:   config.Parity,
	}
}
