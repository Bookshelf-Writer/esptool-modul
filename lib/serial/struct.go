package serial

import (
	"go.bug.st/serial"
	"time"
)

//###########################################################//

type SerialObj struct {
	port        string
	baudRate    int
	readTimeout time.Duration

	dataBits int
	stopBits serial.StopBits
	parity   serial.Parity

	serial serial.Port

	Timeout  ReadTimeoutObj
	BaudRate BaudRateObj
}

func New(port string, baudRate int, dataBits int, stopBits serial.StopBits, parity serial.Parity) (*SerialObj, error) {
	if baudRate < BaudRateMin {
		return nil, ErrBaudRateMin
	}

	obj := SerialObj{}

	obj.port = port
	obj.baudRate = baudRate
	obj.dataBits = dataBits
	obj.stopBits = stopBits
	obj.parity = parity

	obj.Timeout.ss = &obj
	obj.BaudRate.ss = &obj

	return obj.start()
}

func NewEsp(portName string, baudRate int) (*SerialObj, error) {
	return New(portName, baudRate, 8, serial.OneStopBit, serial.NoParity)
}

////

func (ss *SerialObj) mode() *serial.Mode {
	return &serial.Mode{
		BaudRate: ss.baudRate,
		DataBits: ss.dataBits,
		StopBits: ss.stopBits,
		Parity:   ss.parity,
	}
}

func (ss *SerialObj) start() (*SerialObj, error) {
	var err error
	ss.serial, err = serial.Open(ss.port, ss.mode())
	return ss, err
}
