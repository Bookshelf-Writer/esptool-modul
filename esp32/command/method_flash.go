package command

import (
	"github.com/Bookshelf-Writer/esptool-modul/esp32/code"
)

//###########################################################//

type FlashCommandObj struct{}

var Flash FlashCommandObj

////

func (FlashCommandObj) Begin(eraseSize uint32, numBlocks uint32, blockSize uint32, offset uint32) *CommandObj {
	buf := initBuffer()

	buf.Uint32(eraseSize)
	buf.Uint32(numBlocks)
	buf.Uint32(blockSize)
	buf.Uint32(offset)

	return newRequest(code.OpFlashBegin, buf.Bytes())
}

func (FlashCommandObj) BeginDeflate(eraseSize uint32, numBlocks uint32, blockSize uint32, offset uint32) *CommandObj {
	buf := initBuffer()

	buf.Uint32(eraseSize)
	buf.Uint32(numBlocks)
	buf.Uint32(blockSize)
	buf.Uint32(offset)

	return newRequest(code.OpFlashDeflateBegin, buf.Bytes())
}

////

func (FlashCommandObj) Data(data []byte, sequence uint32) *CommandObj {
	buf := initBuffer()

	buf.Uint32(uint32(len(data)))
	buf.Uint32(sequence)
	buf.Uint32(0)
	buf.Uint32(0)
	buf.Write(data)

	return newRequest(code.OpFlashData, buf.Bytes()).Checksum(data)
}

func (FlashCommandObj) DataDeflate(data []byte, sequence uint32) *CommandObj {
	buf := initBuffer()

	buf.Uint32(uint32(len(data)))
	buf.Uint32(sequence)
	buf.Uint32(0)
	buf.Uint32(0)
	buf.Write(data)

	return newRequest(code.OpFlashDeflateLData, buf.Bytes()).Checksum(data)
}

////

func End(reboot bool) *CommandObj {
	buf := initBuffer()

	if reboot {
		buf.Uint32(1)
	} else {
		buf.Uint32(0)
	}

	return newRequest(code.OpFlashEnd, buf.Bytes())
}
