package generator

import (
	"bytes"
	"os"
)

//###########################################################//

type GeneratorObj struct {
	val GeneratorByteValueObj
	str GeneratorStringValueObj

	filename string
	buf      bytes.Buffer
}

func Init(name string, file string) *GeneratorObj {
	obj := GeneratorObj{}

	obj.val.maps = make(map[byte]string)
	obj.val.delim = make(map[byte]bool)
	obj.val.name = name

	obj.str.maps = make(map[string]string)
	obj.str.delim = make(map[string]bool)
	obj.str.name = name

	obj.filename = file

	return &obj
}

/////////////

func (gen *GeneratorObj) Write(data []byte) *GeneratorObj {
	gen.buf.Write(data)
	return gen
}

func (gen *GeneratorObj) Len() int {
	return gen.buf.Len()
}

func (gen *GeneratorObj) SaveFile(pack string) error {
	file, err := os.OpenFile(gen.filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(head(pack))
	file.WriteString("type " + gen.TypeName() + " byte\n\n")

	file.Write(gen.buf.Bytes())
	return nil
}

func (gen *GeneratorObj) SaveFileBuf(pack string) error {
	file, err := os.OpenFile(gen.filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(head(pack))
	file.Write(gen.buf.Bytes())
	return nil
}
