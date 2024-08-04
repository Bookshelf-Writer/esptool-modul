package esp32

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"github.com/Bookshelf-Writer/esptool-modul/esp32/command"
	"time"
)

const blockLengthReadMax uint32 = 64 // TODO check if this value taken from the esptool.py is really true
const blockLengthWriteMax uint32 = 0x400

func (e *ESP32ROM) AttachSpiFlash() (err error) {
	_, err = CheckExecuteCommand(e.SerialPort,
		command.AttachSpiFlash(),
		e.defaultTimeout,
		e.defaultRetries,
	)
	if err != nil {
		return err
	}
	e.flashAttached = true
	e.logger.Print("Attach SPI flash success")
	return
}

func (e *ESP32ROM) ReadFlash(offset uint32, size uint32) ([]byte, error) {
	if !e.flashAttached {
		err := e.AttachSpiFlash()
		if err != nil {
			return []byte{}, err
		}
	}

	receivedData := make([]byte, 0)
	pathTime := time.Now()
	e.log.Info().
		Int("load", len(receivedData)).
		Uint("total", uint(size)).
		Msg(fmt.Sprintf("\t%.2f%s", float64(len(receivedData))/float64(size)*100.0, "%"))

	for {

		if pathTime.Add(time.Second).Before(time.Now()) {
			e.log.Info().
				Int("load", len(receivedData)).
				Uint("total", uint(size)).
				Msg(fmt.Sprintf("\t%.2f%s", float64(len(receivedData))/float64(size)*100.0, "%"))
			pathTime = time.Now()
		}

		if len(receivedData) >= int(size) {
			return receivedData, nil
		}

		blockLength := size - uint32(len(receivedData))
		if blockLength > blockLengthReadMax {
			blockLength = blockLengthReadMax
		}

		response, err := CheckExecuteCommand(e.SerialPort,
			command.Read.Flash(offset+uint32(len(receivedData)), blockLength),
			e.defaultTimeout,
			e.defaultRetries,
		)
		if err != nil {
			return receivedData, err
		}

		receivedData = append(receivedData, response.Data()[:blockLength]...)
	}
}

func compressImage(data []byte) ([]byte, error) {
	var b bytes.Buffer

	w, err := zlib.NewWriterLevel(&b, 9)
	_, err = w.Write(data)
	w.Close()
	return b.Bytes(), err
}

func (e *ESP32ROM) WriteFlash(offset uint32, data []byte, useCompression bool) (err error) {
	if !e.flashAttached {
		err = e.AttachSpiFlash()
		if err != nil {
			return err
		}
	}

	var remaining []byte

	numBlocks := (uint32(len(data)) + blockLengthWriteMax - 1) / blockLengthWriteMax
	e.logger.Print("Start Erase procedure")

	if useCompression {
		remaining, err = compressImage(data)
		if err != nil {
			return err
		}
		uncompressedNumBlocks := numBlocks
		numBlocks = (uint32(len(remaining)) + blockLengthWriteMax - 1) / blockLengthWriteMax
		e.logger.Printf("Compressed %d bytes to %d bytes. Ration = %.1f", len(data), len(remaining), float64(len(remaining))/float64(len(data)))
		_, err = CheckExecuteCommand(e.SerialPort,
			command.Flash.BeginDeflate(
				uncompressedNumBlocks*blockLengthWriteMax,
				numBlocks,
				blockLengthWriteMax,
				offset,
			),
			10*time.Second,
			e.defaultRetries)
	} else {
		remaining = make([]byte, len(data))
		copy(remaining, data)
		_, err = CheckExecuteCommand(e.SerialPort,
			command.Flash.Begin(
				uint32(len(data)),
				numBlocks,
				blockLengthWriteMax,
				offset,
			),
			10*time.Second,
			e.defaultRetries,
		)
	}

	e.logger.Printf("Block size is %d, block count is %d", blockLengthWriteMax, numBlocks)
	if err != nil {
		return err
	}
	e.logger.Print("Begin Flash success.")

	sequence := uint32(0)

	sent := uint32(0)
	total := uint32(len(remaining))

	time.Sleep(20 * time.Millisecond)

	pathTime := time.Now()
	e.log.Info().
		Uint("sent", uint(sent)).
		Uint("total", uint(total)).
		Msg(fmt.Sprintf("\t%.2f%s", float64(sent)/float64(total)*100.0, "%"))

	for {
		if sent >= total {
			break
		}

		if pathTime.Add(time.Second).Before(time.Now()) {
			e.log.Info().
				Uint("sent", uint(sent)).
				Uint("total", uint(total)).
				Msg(fmt.Sprintf("\t%.2f%s", float64(sent)/float64(total)*100.0, "%"))
			pathTime = time.Now()
		}

		blockLength := total - sent
		if blockLength > blockLengthWriteMax {
			blockLength = blockLengthWriteMax
		}
		block := remaining[sent : sent+blockLength]

		if !useCompression && blockLength < blockLengthWriteMax {
			block = append(block, bytes.Repeat([]byte{0xFF}, int(blockLengthWriteMax-blockLength))...)
		}

		for retryCount := 0; retryCount < 3; retryCount++ {
			if retryCount > 0 {
				e.log.Debug().Msg("Received error while writing to Flash")
			}
			if useCompression {
				_, err = CheckExecuteCommand(e.SerialPort,
					command.Flash.DataDeflate(
						block,
						sequence,
					),
					e.defaultTimeout*100,
					e.defaultRetries,
				)
				if err == nil {
					break
				}
			} else {
				_, err = CheckExecuteCommand(e.SerialPort,
					command.Flash.Data(
						block,
						sequence,
					),
					e.defaultTimeout*10,
					e.defaultRetries,
				)

				if err == nil {
					break
				}
			}
		}
		if err != nil {
			return err
		}

		sequence++
		sent += blockLength
	}

	_, err = CheckExecuteCommand(e.SerialPort,
		command.Flash.End(true),
		e.defaultTimeout,
		e.defaultRetries,
	)

	return err
}
