package generator

//###########################################################//

func (gen *GeneratorObj) GetByteVal() *GeneratorByteValueObj {
	return &gen.val
}

func (gen *GeneratorObj) GetStringVal() *GeneratorStringValueObj {
	return &gen.str
}

///////

func (gen *GeneratorObj) GetList() []byte {
	return gen.val.list
}

func (gen *GeneratorObj) GetStrings() []string {
	return gen.str.list
}

///////

func (gen *GeneratorObj) GetText(code byte) string {
	return gen.val.maps[code]
}

func (gen *GeneratorObj) GetTextString(code string) string {
	return gen.str.maps[code]
}

///////

func (gen *GeneratorObj) SetDelim(code byte) {
	_, ok := gen.val.delim[code]
	if ok {
		gen.LN()
	}
}

func (gen *GeneratorObj) SetDelimString(code string) {
	_, ok := gen.str.delim[code]
	if ok {
		gen.LN()
	}
}
