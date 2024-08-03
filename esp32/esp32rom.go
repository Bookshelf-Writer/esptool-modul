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

func (e *ESP32ROM) Reset() (err error) {
	// set IO0=HIGH
	err = e.SerialPort.SetDTR(false)
	if err != nil {
		return
	}
	// set EN=LOW, chip in reset
	err = e.SerialPort.SetRTS(true)
	if err != nil {
		return
	}

	time.Sleep(100 * time.Millisecond)

	// set IO0=LOW
	err = e.SerialPort.SetDTR(true)
	if err != nil {
		return
	}
	// EN=HIGH, chip out of reset
	err = e.SerialPort.SetRTS(false)

	time.Sleep(5 * time.Millisecond)
	return
}

func (e *ESP32ROM) Connect(maxRetries uint) (err error) {
	err = e.Reset()
	if err != nil {
		return
	}

	err = e.SerialPort.Flush()
	if err != nil {
		return
	}

	for i := uint(0); i < maxRetries; i++ {
		e.log.Debug().Msg("Connecting")
		err = e.Sync()
		if err == nil {
			break
		}
	}
	return
}

func (e *ESP32ROM) Sync() (err error) {
	response, err := e.ExecuteCommand(
		command.Sync(),
		1000*time.Millisecond,
	)
	if err != nil {
		return err
	}
	if response.Status != true {
		err = fmt.Errorf("Command failed")
	}
	return
}

func (e *ESP32ROM) ReadEfuse(efuseIndex uint) ([4]byte, error) {
	return e.ReadRegister(efuseRegBase + (4 * efuseIndex))
}

func (e *ESP32ROM) ReadRegister(register uint) ([4]byte, error) {
	response, err := e.ExecuteCommand(
		command.Read.Register(uint32(register)),
		e.defaultTimeout,
	)
	if err != nil {
		return [4]byte{}, err
	}
	return [4]byte(response.Checksum()), nil
}

func (e *ESP32ROM) ExecuteCommand(command *command.CommandObj, timeout time.Duration) (*portal.ResponseObj, error) {
	err := portal.Write(e.SerialPort, command.Bytes())
	if err != nil {
		return nil, err
	}
	for retryCount := 0; retryCount < 16; retryCount++ {

		responseBuf, err := portal.Read(e.SerialPort, timeout)
		if err != nil {
			return nil, err
		}
		if responseBuf[1] != command.OpcodeToByte() {
			e.log.Trace().Msg("Opcode did not match")
			continue
		} else {
			return portal.Response(responseBuf)
		}
	}
	return nil, fmt.Errorf("Retrycount exceeded")
}

func (e *ESP32ROM) CheckExecuteCommand(command *command.CommandObj, timeout time.Duration, retries int) (response *portal.ResponseObj, err error) {
	for retryCount := 0; retryCount < retries; retryCount++ {
		response, err = e.ExecuteCommand(command, timeout)
		if err != nil {
			e.log.Debug().Str("command", command.Opcode()).Msg("Executing command failed. Retrying")
			continue
		}
		if !response.Status {
			err = fmt.Errorf("Device returned for command %s status %s", command.Opcode(), response.String())
			e.log.Debug().Str("command", command.Opcode()).Msg("Received non success status for command. Retrying")
			continue
		} else {
			break
		}
	}
	return
}

func (e *ESP32ROM) ChangeBaudrate(newBaudrate uint32) error {
	e.log.Trace().Uint("rate", uint(newBaudrate)).Msg("Changing BaudRate")
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

	e.log.Debug().Uint("rate", uint(e.SerialPort.BaudRate.Get())).Msg("Changed BaudRate")
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
