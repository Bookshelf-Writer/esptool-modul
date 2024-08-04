package address

// ###########################################################//

type AddressType uint64

type AddressCellObj struct {
	Begin AddressType
	End   AddressType
}

func (adr *AddressCellObj) Size() AddressType {
	return adr.End - adr.Begin
}
