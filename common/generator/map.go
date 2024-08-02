package generator

//###########################################################//

type GeneratorByteValueObj struct {
	maps    map[byte]string
	delim   map[byte]bool
	list    []byte
	lastKey byte

	name string
}

func (arr *GeneratorByteValueObj) Add(code byte, text string) *GeneratorByteValueObj {
	arr.lastKey = code

	arr.list = append(arr.list, code)
	arr.maps[code] = text

	return arr
}

func (arr *GeneratorByteValueObj) Delim() {
	arr.delim[arr.lastKey] = true
}

///////////////////////////////////////////

type GeneratorStringValueObj struct {
	maps    map[string]string
	delim   map[string]bool
	list    []string
	lastKey string

	name string
}

func (arr *GeneratorStringValueObj) Add(code string, text string) *GeneratorStringValueObj {
	arr.lastKey = code

	arr.list = append(arr.list, code)
	arr.maps[code] = text

	return arr
}

func (arr *GeneratorStringValueObj) Delim() {
	arr.delim[arr.lastKey] = true
}
