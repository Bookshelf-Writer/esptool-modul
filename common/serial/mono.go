package serial

import (
	"github.com/jacobsa/go-serial/serial"
	"io"
	"time"
)

type StopBits byte

const (
	StopBitsOne StopBits = iota
	StopBitsTwo
)

type Parity byte

const (
	ParityNone Parity = iota
	ParityOdd
	ParityEven
)

type Config struct {
	PortName    string
	BaudRate    uint32
	ReadTimeout time.Duration
	DataBits    uint
	StopBits    StopBits
	Parity      Parity
}

func NewConfig(portName string, baudRate uint32) *Config {
	return &Config{
		PortName:    portName,
		BaudRate:    baudRate,
		ReadTimeout: time.Millisecond,
		DataBits:    8,
		StopBits:    StopBitsOne,
		Parity:      ParityNone,
	}
}

type Port struct {
	port io.ReadWriteCloser
}

func OpenPort(config *Config) (*Port, error) {
	options := serial.OpenOptions{
		PortName:              config.PortName,
		BaudRate:              uint(config.BaudRate),
		DataBits:              config.DataBits,
		StopBits:              uint(config.StopBits),
		ParityMode:            serial.ParityMode(config.Parity),
		InterCharacterTimeout: uint(config.ReadTimeout.Milliseconds()),
		MinimumReadSize:       1,
	}

	port, err := serial.Open(options)
	if err != nil {
		return nil, err
	}

	return &Port{port: port}, nil
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
	return nil
}

func (p *Port) SetRTS(rts bool) error {
	return nil
}
