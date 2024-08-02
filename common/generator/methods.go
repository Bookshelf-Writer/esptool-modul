package generator

import (
	"fmt"
	"strings"
)

//###########################################################//

func (gen *GeneratorObj) LN() *GeneratorObj {
	gen.Write([]byte("\n"))
	return gen
}

func (gen *GeneratorObj) Byte(b byte) *GeneratorObj {
	gen.Write([]byte(fmt.Sprintf("%d", b)))
	return gen
}

func (gen *GeneratorObj) Print(text string) *GeneratorObj {
	gen.Write([]byte(text))
	return gen
}

func (gen *GeneratorObj) PrintString(text string) *GeneratorObj {
	gen.Write([]byte("\"" + text + "\""))
	return gen
}

func (gen *GeneratorObj) PrintLN(text string) *GeneratorObj {
	gen.Print(text).LN()
	return gen
}

func (gen *GeneratorObj) Del(len int) *GeneratorObj {
	gen.buf.Truncate(gen.Len() - len)
	return gen
}

func (gen *GeneratorObj) Repeat(pos int) *GeneratorObj {
	gen.buf.Write([]byte(strings.Repeat("\t", pos)))
	return gen
}

////

func (gen *GeneratorObj) TypeName() string {
	return gen.val.name + "Type"
}

func (gen *GeneratorObj) Type() *GeneratorObj {
	gen.Write([]byte(gen.TypeName()))
	return gen
}

func (gen *GeneratorObj) Map() *GeneratorObj {
	gen.Write([]byte(gen.val.name + "Map"))
	return gen
}

func (gen *GeneratorObj) TitleCase(text string) *GeneratorObj {
	gen.Write([]byte(toTitleCase(text)))
	return gen
}

func (gen *GeneratorObj) ConstCode(code string) *GeneratorObj {
	gen.Write([]byte(gen.val.name + toTitleCase(code)))
	return gen
}

func (gen *GeneratorObj) ConstText(code string) *GeneratorObj {
	gen.Write([]byte(gen.val.name + "Text" + toTitleCase(code)))
	return gen
}
