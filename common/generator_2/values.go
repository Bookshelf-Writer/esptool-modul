package generator_2

//###########################################################//

type GeneratorValueObj struct {
	maps    map[any]string
	delim   map[any]bool
	list    []any
	lastKey any

	Get GeneratorValueGetObj

	name *string
}

func generatorValueInit(name *string) *GeneratorValueObj {
	obj := GeneratorValueObj{}

	obj.maps = make(map[any]string)
	obj.delim = make(map[any]bool)

	obj.name = name
	obj.Get.arr = &obj

	return &obj
}

func (arr *GeneratorValueObj) Add(code any, text string) *GeneratorValueObj {
	arr.lastKey = code

	arr.list = append(arr.list, code)
	arr.maps[code] = text

	return arr
}

func (arr *GeneratorValueObj) Delim() {
	arr.delim[arr.lastKey] = true
}

////

type GeneratorValueGetObj struct {
	arr *GeneratorValueObj
}

func (get *GeneratorValueGetObj) Text(code any) string {
	return get.arr.maps[code]
}

func (get *GeneratorValueGetObj) IsDelim(code any) bool {
	_, ok := get.arr.delim[code]
	return ok
}

//

func (get *GeneratorValueGetObj) Bytes() []byte {
	bufArr := make([]byte, len(get.arr.list))
	for pos, val := range get.arr.list {
		bufArr[pos] = val.(byte)
	}
	return bufArr
}

func (get *GeneratorValueGetObj) Strings() []string {
	bufArr := make([]string, len(get.arr.list))
	for pos, val := range get.arr.list {
		bufArr[pos] = val.(string)
	}
	return bufArr
}
