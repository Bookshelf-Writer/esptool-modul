package command

//###########################################################//

type NumberObj struct{}

var Number NumberObj

////

func (NumberObj) Uint16(data uint16) []byte {
	return initBuffer().Uint16(data).Bytes()
}

func (NumberObj) Uint32(data uint32) []byte {
	return initBuffer().Uint32(data).Bytes()
}
