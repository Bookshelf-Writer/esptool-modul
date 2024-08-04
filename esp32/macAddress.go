package esp32

import (
	"net"
)

func (e *ESP32ROM) GetChipMAC() (string, error) {
	data, err := GetUID(e.SerialPort, e.defaultTimeout)
	if err != nil {
		return "", err
	}

	macBuf := make([]byte, 6)
	macBuf[0] = data[1]
	macBuf[1] = data[0]
	macBuf[2] = data[7]
	macBuf[3] = data[6]
	macBuf[4] = data[5]
	macBuf[5] = data[4]

	mac := net.HardwareAddr(macBuf)
	return mac.String(), nil
}
