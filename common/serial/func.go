package serial

import "time"

//###########################################################//

func (p *PortObj) Flush() error {
	err := p.port.ResetInputBuffer()
	if err != nil {
		return err
	}

	return p.port.ResetInputBuffer()
}

////////

func (p *PortObj) Reset() error {
	// set IO0=HIGH
	err := p.port.SetDTR(false)
	if err != nil {
		return err
	}

	// set EN=LOW, chip in reset
	err = p.port.SetRTS(true)
	if err != nil {
		return err
	}

	time.Sleep(100 * time.Millisecond)

	// set IO0=LOW
	err = p.port.SetDTR(true)
	if err != nil {
		return err
	}

	// EN=HIGH, chip out of reset
	err = p.port.SetRTS(false)

	time.Sleep(5 * time.Millisecond)
	return nil
}

func (p *PortObj) Connect() error {
	err := p.Reset()
	if err != nil {
		return err
	}

	err = p.Flush()
	if err != nil {
		return err
	}

	return nil
}
