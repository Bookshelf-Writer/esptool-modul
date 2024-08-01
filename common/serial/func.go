package serial

//###########################################################//

func (p *PortObj) Flush() error {
	err := p.port.ResetInputBuffer()
	if err != nil {
		return err
	}

	return p.port.ResetInputBuffer()
}

////////

func (br *PortBaudRateObj) Set(newBaudrate uint32) error {
	br.obj.conf.BaudRate = newBaudrate
	return br.obj.port.SetMode(br.obj.conf.Mode())
}

func (br *PortBaudRateObj) Get() uint32 {
	return br.obj.conf.BaudRate
}
