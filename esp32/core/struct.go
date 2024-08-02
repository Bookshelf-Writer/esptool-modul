package core

//###########################################################//

type ModulMapStruct struct {
	Start uint64 `json:"start"`
	End   uint64 `json:"end"`
}

type ModulEncryptStruct struct {
	WriteAlign int  `json:"write_align"`
	Supports   bool `json:"supports"`
}

type ModulSystemStruct struct {
	Chip      int `json:"chip"`
	LenStatus int `json:"len_status"`

	FlashOffset uint64 `json:"flash_offset"`
	UF2         uint64 `json:"uf2"`
}

type ModulStruct struct {
	Name string `json:"name"`

	Sys     ModulSystemStruct `json:"sys"`
	Encrypt ModulEncryptStruct

	MagicValue []uint64                  `json:"magicValue"`
	Memory     map[string]ModulMapStruct `json:"memory"`

	Value *map[string]string `json:"value,omitempty"`
}
