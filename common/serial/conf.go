package serial

import (
	"go.bug.st/serial"
	"time"
)

//###########################################################//

type ConfigObj struct {
	PortName    string
	BaudRate    uint32
	ReadTimeout time.Duration
	DataBits    int
	StopBits    serial.StopBits
	Parity      serial.Parity
}

func NewConfig(portName string, baudRate uint32) *ConfigObj {
	return &ConfigObj{
		PortName:    portName,
		BaudRate:    baudRate,
		ReadTimeout: time.Millisecond,
		DataBits:    8,
		StopBits:    serial.OneStopBit,
		Parity:      serial.NoParity,
	}
}
