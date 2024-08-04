package esp32

import (
	"bytes"
	"fmt"
	"github.com/Bookshelf-Writer/esptool-modul/common/output"
	"github.com/Bookshelf-Writer/esptool-modul/common/serial"
	"github.com/Bookshelf-Writer/esptool-modul/esp32/command"
	"github.com/Bookshelf-Writer/esptool-modul/esp32/portal"
	"github.com/rs/zerolog"
	"time"
)

const (
	efuseRegBase    uint = 0x6001a000
	drRegSysconBase uint = 0x3ff66000
	macEfuseReg     uint = 0x3f41A044 // ESP32-S2 has special block for MAC efuses
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

func (e *ESP32ROM) CheckExecuteCommand(command *command.CommandObj, timeout time.Duration, retries int) (*portal.ResponseObj, error) {
	return CheckExecuteCommand(e.SerialPort, command, timeout, retries)
}

func (e *ESP32ROM) ChangeBaudrate(newBaudrate uint32) error {
	_, err := e.CheckExecuteCommand(
		command.ChangeBaudRate(newBaudrate, 0),
		e.defaultTimeout,
		e.defaultRetries,
	)
	if err != nil {
		return err
	}

	err = e.SerialPort.BaudRate.Set(newBaudrate)
	if err != nil {
		return err
	}

	e.log.Trace().Uint32("rate", e.SerialPort.BaudRate.Get()).Msg("Changed BaudRate")
	time.Sleep(10 * time.Millisecond)

	e.SerialPort.Flush() // get rid of crap sent during baud rate change
	return nil
}

func (e *ESP32ROM) ReadPartitionList() (PartitionList, error) {
	e.log.Debug().Msg("Reading partiton table")

	bindata, err := e.ReadFlash(uint32(partitionTableOffset), uint32(partitionTableMaxSize))

	if err != nil {
		return PartitionList{}, fmt.Errorf("Could not read partition table from chip: %v", err)
	}

	reader := NewPartitionBinaryReader(bytes.NewReader(bindata))

	return reader.ReadAll()
}
