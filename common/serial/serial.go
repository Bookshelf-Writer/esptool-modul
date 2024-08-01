package serial

import (
	"go.bug.st/serial"
	"time"
)

//###########################################################//

type PortBaudRateObj struct {
	obj *PortObj
}

type PortObj struct {
	port serial.Port
	conf *ConfigObj

	BaudRate PortBaudRateObj
}

func PortInit(config *ConfigObj) (obj *PortObj, err error) {
	obj.conf = config
	obj.BaudRate.obj = obj

	obj.port, err = serial.Open(config.PortName, config.Mode())
	if err != nil {
		return nil, err
	}

	return obj, nil
}

////////

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
