package serial

//###########################################################//

type PortBaudRateObj struct {
	obj *PortObj
}

func (br *PortBaudRateObj) Set(newBaudRate uint32) error {
	br.obj.conf.BaudRate = newBaudRate
	return br.obj.port.SetMode(br.obj.conf.Mode())
}

func (br *PortBaudRateObj) Get() uint32 {
	return br.obj.conf.BaudRate
}
