package common

import (
	"github.com/Bookshelf-Writer/esptool-modul/esp32/command"
)

func NewReadRegisterCommand(register uint32) *command.CommandObj {
	return command.Read.Register(register)
}

func NewSyncCommand() *command.CommandObj {
	return command.Sync()
}

func NewAttachSpiFlashCommand() *command.CommandObj {
	return command.AttachSpiFlash()
}

func NewReadFlashCommand(offset uint32, size uint32) *command.CommandObj {
	return command.Read.Flash(offset, size)
}

func NewChangeBaudrateCommand(newBaudrate uint32, oldBaudrate uint32) *command.CommandObj {
	return command.ChangeBaudRate(newBaudrate, oldBaudrate)
}

func NewBeginFlashCommand(eraseSize uint32, numBlocks uint32, blockSize uint32, offset uint32) *command.CommandObj {
	return command.Flash.Begin(eraseSize, numBlocks, blockSize, offset)
}

func NewBeginFlashDeflCommand(eraseSize uint32, numBlocks uint32, blockSize uint32, offset uint32) *command.CommandObj {
	return command.Flash.BeginDeflate(eraseSize, numBlocks, blockSize, offset)
}

func NewFlashDataCommand(data []byte, sequence uint32) *command.CommandObj {
	return command.Flash.Data(data, sequence)
}

func NewFlashDataDeflCommand(data []byte, sequence uint32) *command.CommandObj {
	return command.Flash.DataDeflate(data, sequence)
}

func NewFlashEndCommand(reboot bool) *command.CommandObj {
	return command.Flash.End(reboot)
}
