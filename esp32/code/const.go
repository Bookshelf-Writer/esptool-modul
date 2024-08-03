package code

//###########################################################//

const (
	MaxRamBlock = uint64(0x1800) //Maximum block sized for RAM and Flash writes, respectively.
	WriteSize   = uint64(0x400)

	ByteImageMagic    = byte(0xE9) //First byte of the application image
	ByteChecksumMagic = byte(0xEF) //Initial state for the checksum routine
)
