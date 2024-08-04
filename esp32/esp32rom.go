package esp32

import (
	"bytes"
	"fmt"
	"github.com/Bookshelf-Writer/esptool-modul/common/serial"
	"github.com/Bookshelf-Writer/esptool-modul/esp32/command"
	"github.com/Bookshelf-Writer/esptool-modul/lib/output"
	"github.com/rs/zerolog"
	"time"
)

type ESP32ROM struct {
	SerialPort     *serial.PortObj
	flashAttached  bool
	logger         *zerolog.Logger
	defaultTimeout time.Duration
	defaultRetries int
	log            output.LogObj
}

func NewESP32ROM(serialPort *serial.PortObj, logger *output.LogObj) *ESP32ROM {
	logger = logger.NewLog("NewESP32ROM")

	return &ESP32ROM{
		SerialPort:     serialPort,
		logger:         logger.ZeroLog(),
		defaultTimeout: 100 * time.Millisecond,
		defaultRetries: 3,
		log:            *logger,
	}
}

func (e *ESP32ROM) Reset() error {
	return e.SerialPort.Reset()
}

////

func (e *ESP32ROM) Sync() (err error) {
	response, err := RunCommand(e.SerialPort, command.Sync(), 400*time.Millisecond)
	if err != nil {
		return err
	}

	if response.Status != true {
		err = fmt.Errorf("Command failed")
	}
	return
}

func (e *ESP32ROM) Connect(maxRetries uint) error {
	err := e.SerialPort.Connect()
	if err != nil {
		return err
	}

	for i := uint(0); i < maxRetries; i++ {
		err = e.Sync()
		if err == nil {
			return nil
		} else {
			e.log.Trace().Err(err).Uint("retry", i).Msg("Connecting")
		}
	}

	return err
}

////

func (e *ESP32ROM) ReadPartitionList() (PartitionList, error) {
	e.log.Debug().Msg("Reading partiton table")

	bindata, err := e.ReadFlash(uint32(partitionTableOffset), uint32(partitionTableMaxSize))

	if err != nil {
		return PartitionList{}, fmt.Errorf("Could not read partition table from chip: %v", err)
	}

	reader := NewPartitionBinaryReader(bytes.NewReader(bindata))

	return reader.ReadAll()
}
