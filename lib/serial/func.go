package serial

//###########################################################//

func Check(portName string) bool {
	serialPort, err := NewEsp(portName, 115200)
	if err != nil {
		return false
	}

	serialPort.Close()
	return true
}
