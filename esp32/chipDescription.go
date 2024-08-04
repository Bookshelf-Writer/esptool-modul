package esp32

import (
	"fmt"
	"github.com/Bookshelf-Writer/esptool-modul/esp32/code"
)

type ChipDescription struct {
	ChipType code.EspType
	Revision byte
}

func (c *ChipDescription) String() string {
	return fmt.Sprintf("%s (revision %d)", c.ChipType.String(), c.Revision)
}

func (e *ESP32ROM) GetChipDescription() (*ChipDescription, error) {
	word3, err := ReadEfuse(e.SerialPort, e.defaultTimeout, 3)
	if err != nil {
		return nil, err
	}

	word5, err := ReadEfuse(e.SerialPort, e.defaultTimeout, 5)
	if err != nil {
		return nil, err
	}

	apbCtlBase, err := ReadRegister(e.SerialPort, e.defaultTimeout, uint32(0x3ff66000+0x7C))

	revisionBit0 := (word3[1] >> 7) & 0x01
	revisionBit1 := (word5[2] >> 4) & 0x01
	revisionBit2 := (apbCtlBase[3] >> 7) & 0x01

	revision := byte(0)
	if revisionBit0 > 0 {
		if revisionBit1 > 0 {
			if revisionBit2 > 0 {
				revision = 3
			} else {
				revision = 2
			}
		} else {
			revision = 1
		}
	}

	return &ChipDescription{
		ChipType: code.EspType((word3[1] >> 1) & 0x07),
		Revision: revision,
	}, nil
}
