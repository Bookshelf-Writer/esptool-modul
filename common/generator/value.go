package generator

//###########################################################//

func (gen *GeneratorObj) GetVal() *GeneratorValueObj {
	return &gen.val
}

func (gen *GeneratorObj) GetList() []byte {
	return gen.val.list
}

func (gen *GeneratorObj) GetText(code byte) string {
	return gen.val.maps[code]
}

func (gen *GeneratorObj) SetDelim(code byte) {
	_, ok := gen.val.delim[code]
	if ok {
		gen.LN()
	}
}
