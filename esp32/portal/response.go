package portal

import (
	"fmt"
	"github.com/Bookshelf-Writer/esptool-modul/esp32/code"
)

//###########################################################//

type ResponseObj struct {
	Status    bool
	ErrorCode code.ErrType

	msg MsgObj
}

func Response(data []byte) (*ResponseObj, error) {
	if len(data) < 10 {
		return nil, ErrResponseLength
	}
	obj := ResponseObj{}

	obj.msg.Direction = code.DirectionType(data[0])
	obj.msg.Opcode = code.OpType(data[1])

	obj.msg.Checksum = data[4:8]
	obj.msg.Data = data[8 : len(data)-1]
	obj.msg.Length = len(obj.msg.Data)

	var err error
	obj.Status, obj.ErrorCode, err = parseStatus(data[len(data)-2:])

	return &obj, err
}

func parseStatus(data []byte) (bool, code.ErrType, error) {
	if len(data) != 2 {
		return false, 0, ErrResponseStatusLength
	}

	return data[0] == 0, code.ErrType(data[1]), nil
}

////

func (r *ResponseObj) String() string {
	return fmt.Sprintf("success = %t, error_code = %s", r.Status, r.ErrorCode.String())
}

func (r *ResponseObj) Checksum() []byte {
	return r.msg.Checksum
}

func (r *ResponseObj) Data() []byte {
	return r.msg.Data
}
