package serial

import "time"

//###########################################################//

func (ss *SerialObj) Close() error {
	return ss.serial.Close()
}

func (ss *SerialObj) Write(b []byte) (int, error) {
	return ss.serial.Write(b)
}

func (ss *SerialObj) Read(b []byte) (int, error) {
	return ss.serial.Read(b)
}

////

func (ss *SerialObj) Flush() error {
	err := ss.serial.ResetOutputBuffer()
	if err != nil {
		return err
	}

	return ss.serial.ResetInputBuffer()
}

func (ss *SerialObj) Reset() error {
	err := ss.serial.SetDTR(false) // set IO0=HIGH
	if err != nil {
		return err
	}

	err = ss.serial.SetRTS(true) // set EN=LOW, chip in reset
	if err != nil {
		return err
	}

	time.Sleep(100 * time.Millisecond)

	err = ss.serial.SetDTR(true) // set IO0=LOW
	if err != nil {
		return err
	}

	err = ss.serial.SetRTS(false) // EN=HIGH, chip out of reset
	if err != nil {
		return err
	}

	time.Sleep(5 * time.Millisecond)
	return nil
}
