package common

import (
	"github.com/Bookshelf-Writer/esptool-modul/common/serial"
	"github.com/Bookshelf-Writer/esptool-modul/esp32/portal"
	"github.com/rs/zerolog"
	"time"
)

type SlipReadWriter struct {
	BaseReadWriter *serial.PortObj
	Timeout        time.Duration
	logger         *zerolog.Logger
}

func NewSlipReadWriter(base *serial.PortObj, logger *zerolog.Logger) *SlipReadWriter {
	return &SlipReadWriter{
		BaseReadWriter: base,
		logger:         logger,
	}
}

func (s *SlipReadWriter) Write(b []byte) error {
	return portal.Write(s.BaseReadWriter, b)
}

func (s *SlipReadWriter) Read(timeout time.Duration) ([]byte, error) {
	return portal.Read(s.BaseReadWriter, timeout)
}
