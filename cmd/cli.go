package main

import (
	"flag"
	"time"
)

//###########################################################//

type CliBaudObj struct {
	Connect  *uint
	Transfer *uint
}

type CliConnObj struct {
	Timeout *time.Duration
	Retries *uint
}

type CliFlashObj struct {
	Compress *bool
	Offset   *uint64
	Size     *uint64
	FilePath *string
}

////

type CliObj struct {
	Port *string

	Baud  CliBaudObj
	Conn  CliConnObj
	Flash CliFlashObj
}

var CLI = CliObj{}

////

func init() {
	CLI.Port = flag.String(CliValPort, "", CliValTextPort)
	CLI.Baud.Connect = flag.Uint(CliValBaudConnect, 115200, CliValTextBaudConnect)
	CLI.Baud.Transfer = flag.Uint(CliValBaudTransfer, 921600, CliValTextBaudTransfer)

	CLI.Conn.Timeout = flag.Duration(CliValConnTimeout, 800*time.Millisecond, CliValTextConnTimeout)
	CLI.Conn.Retries = flag.Uint(CliValConnRetries, 8, CliValTextConnRetries)

	CLI.Flash.Compress = flag.Bool(CliValFlashCompress, false, CliValTextFlashCompress)
	CLI.Flash.Offset = flag.Uint64(CliValFlashOffset, 0, CliValTextFlashOffset)
	CLI.Flash.Size = flag.Uint64(CliValFlashSize, 0, CliValTextFlashSize)
	CLI.Flash.FilePath = flag.String(CliValFlashFile, "", CliValTextFlashFile)
}
