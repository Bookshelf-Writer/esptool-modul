package common

import "encoding/binary"

//###########################################################//

func Uint32ToBytes(value uint32) []byte {
	var buf []byte
	binary.BigEndian.PutUint32(buf, value)
	return buf
}
