package portal

import (
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
		return ErrMismatchBytes
	}

	return nil
}
