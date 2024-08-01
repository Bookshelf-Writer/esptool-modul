package serial

import (
	"go.bug.st/serial"
	"time"
)

//###########################################################//

type PortObj struct {
	port serial.Port
	mode *serial.Mode
}

func OpenPort(config *ConfigObj) (*PortObj, error) {
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

	return &PortObj{port: port, mode: mode}, nil
}

func (p *PortObj) Close() error {
	return p.port.Close()
}

func (p *PortObj) Write(b []byte) (int, error) {
	return p.port.Write(b)
}

func (p *PortObj) Read(b []byte) (int, error) {
	return p.port.Read(b)
}

func (p *PortObj) SetDTR(dtr bool) error {
	return p.port.SetDTR(dtr)
}

func (p *PortObj) SetRTS(rts bool) error {
	return p.port.SetRTS(rts)
}

func (p *PortObj) SetReadTimeout(t time.Duration) error {
	return p.port.SetReadTimeout(t)
}

func (p *PortObj) Flush() error {
	err := p.port.ResetInputBuffer()
	if err != nil {
		return err
	}

	return p.port.ResetInputBuffer()
}

func (p *PortObj) SetBaudrate(newBaudrate uint32) error {
	return p.port.SetMode(&serial.Mode{
		BaudRate: int(newBaudrate),
		DataBits: p.mode.DataBits,
		StopBits: p.mode.StopBits,
		Parity:   p.mode.Parity,
	})
}

func (p *PortObj) GetBaudrate() int {
	return p.mode.BaudRate
}
