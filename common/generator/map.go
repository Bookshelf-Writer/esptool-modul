package generator

//###########################################################//

type GeneratorValueObj struct {
	maps    map[byte]string
	delim   map[byte]bool
	list    []byte
	lastKey byte

	name string
}

func (arr *GeneratorValueObj) Add(code byte, text string) *GeneratorValueObj {
	arr.lastKey = code

	arr.list = append(arr.list, code)
	arr.maps[code] = text

	return arr
}

func (arr *GeneratorValueObj) Delim() {
	arr.delim[arr.lastKey] = true
}
