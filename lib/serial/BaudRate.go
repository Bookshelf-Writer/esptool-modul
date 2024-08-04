package serial

//###########################################################//

type BaudRateObj struct {
	ss *SerialObj
}

func (obj *BaudRateObj) Set(newBaudRate int) error {
	if newBaudRate < BaudRateMin {
		return ErrBaudRateMin
	}

	obj.ss.baudRate = newBaudRate
	return obj.ss.serial.SetMode(obj.ss.mode())
}

func (obj *BaudRateObj) Get() int {
	return obj.ss.baudRate
}
