package portal

import (
	"fmt"
	"github.com/Bookshelf-Writer/esptool-modul/common/serial"
)

//###########################################################//

func Write(port *serial.PortObj, data []byte) error {
	data = encode(data)

	n, err := port.Write(data)
	if err != nil {
		return err
	}

	if n != len(data) {
		err = fmt.Errorf("Expected to send %d bytes but transfered only %d bytes.", len(data), n)
		return err
	}

	return nil
}
