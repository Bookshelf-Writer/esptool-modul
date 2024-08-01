package serial

import (
	"go.bug.st/serial"
	"time"
)

type Config struct {
	PortName    string
	BaudRate    uint32
	ReadTimeout time.Duration
	DataBits    int
	StopBits    serial.StopBits
	Parity      serial.Parity
}

func NewConfig(portName string, baudRate uint32) *Config {
	return &Config{
		PortName:    portName,
		BaudRate:    baudRate,
		ReadTimeout: time.Millisecond,
		DataBits:    8,
		StopBits:    serial.OneStopBit,
		Parity:      serial.NoParity,
	}
}

type Port struct {
	port serial.Port
	mode *serial.Mode
}

func OpenPort(config *Config) (*Port, error) {
	mode := &serial.Mode{
		BaudRate: int(config.BaudRate),
		DataBits: config.DataBits,
		StopBits: config.StopBits,
		Parity:   config.Parity,
	}

	port, err := serial.Open(config.PortName, mode)
	if err != nil {
		return nil, err
	}

	return &Port{port: port, mode: mode}, nil
}

func (p *Port) Close() error {
	return p.port.Close()
}

func (p *Port) Write(b []byte) (int, error) {
	return p.port.Write(b)
}

func (p *Port) Read(b []byte) (int, error) {
	return p.port.Read(b)
}

func (p *Port) SetDTR(dtr bool) error {
	return p.port.SetDTR(dtr)
}

func (p *Port) SetRTS(rts bool) error {
	return p.port.SetRTS(rts)
}

func (p *Port) SetReadTimeout(t time.Duration) error {
	return p.port.SetReadTimeout(t)
}

func (p *Port) Flush() error {
	err := p.port.ResetInputBuffer()
	if err != nil {
		return err
	}

	return p.port.ResetInputBuffer()
}

func (p *Port) SetBaudrate(newBaudrate uint32) error {
	return p.port.SetMode(&serial.Mode{
		BaudRate: int(newBaudrate),
		DataBits: p.mode.DataBits,
		StopBits: p.mode.StopBits,
		Parity:   p.mode.Parity,
	})
}

func (p *Port) GetBaudrate() int {
	return p.mode.BaudRate
}
