package main

import (
	"encoding/csv"
	"github.com/Bookshelf-Writer/esptool-modul/lib/generator"
	"os"
	"strings"
)

//###########################################################//

func readCSV(filename string) [][]string {
	file, err := os.Open("esp32/address/tables/" + filename + ".csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var rows [][]string

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	// Читаємо файли поки не дійдемо до кінця.
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		var buf []string
		for _, cell := range record {
			buf = append(buf, strings.TrimSpace(cell))
		}
		rows = append(rows, buf)
	}

	return rows
}

type adrObj struct {
	begin   string
	end     string
	name    string
	comment string
}

func build(
	obj *generator.GeneratorObj,
	valName *generator.GeneratorValueObj,
	valBegin *generator.GeneratorValueObj,
	valEnd *generator.GeneratorValueObj,
	valComment *generator.GeneratorValueObj,
) {

	obj.PrintLN("const (")
	for _, code := range valName.Get.Ints() {
		text := valName.Get.Text(code)

		obj.Offset(1).Name.CodeToTitleCase(text).Print(" ")
		obj.Name.Type().Print(" = ").Number(code)
		obj.Offset(1).Print("//" + valComment.Get.Text(code))
		obj.LN()
	}
	obj.PrintLN(")").LN()

	//

	obj.PrintLN("const (")
	for _, code := range valName.Get.Ints() {
		text := valName.Get.Text(code)

		obj.Offset(1).Name.SelfParamCode("AdrBegin", text)
		obj.Print(" AddressType = ")
		obj.Print(valBegin.Get.Text(code)).LN()
	}
	obj.PrintLN(")").LN()

	obj.PrintLN("const (")
	for _, code := range valName.Get.Ints() {
		text := valName.Get.Text(code)

		obj.Offset(1).Name.SelfParamCode("AdrEnd", text)
		obj.Print(" AddressType = ")
		obj.Print(valEnd.Get.Text(code)).LN()
	}
	obj.PrintLN(")").LN()

	obj.PrintLN("var (")
	for _, code := range valName.Get.Ints() {
		text := valName.Get.Text(code)

		obj.Offset(1).Name.SelfParamCode("Adr", text)
		obj.Print(" = AddressCellObj{")
		obj.Name.SelfParamCode("AdrBegin", text).Print(",")
		obj.Name.SelfParamCode("AdrEnd", text).Print("}")
		obj.Offset(1).Print("//" + valComment.Get.Text(code))
		obj.LN()
	}
	obj.PrintLN(")").LN()

	//

	obj.Print("var ").Name.Map().Print(" = map[").Name.Type().PrintLN("]*AddressCellObj{")
	for _, code := range valName.Get.Ints() {
		text := valName.Get.Text(code)

		obj.Offset(1).Name.CodeToTitleCase(text).Print(": &")
		obj.Name.SelfParamCode("Adr", text).PrintLN(",")

	}
	obj.PrintLN("}").LN()

	//
	file := obj.Save("address")
	file.Add.Type(obj.Name.GetType(), "byte")
	err := file.End()
	if err != nil {
		panic(err)
	}
}

//###########################################################//

func main() {
	peripheral()
	syscon()
	DPort()
}

////

func peripheral() {
	name := "peripheral"

	bufArr := readCSV(name)

	var valueArr []adrObj
	for _, buf := range bufArr {
		if len(buf[0]) > 0 {
			obj := adrObj{
				buf[1],
				buf[2],
				buf[4],
				buf[3],
			}
			if len(buf) == 6 {
				obj.comment = strings.TrimSpace(obj.comment + "\t\t" + buf[5])
			}
			valueArr = append(valueArr, obj)
		}
	}

	//

	obj := generator.Init(name, "esp32/address/"+name+".go")

	valName := obj.NewVal()
	valBegin := obj.NewVal()
	valEnd := obj.NewVal()
	valComment := obj.NewVal()

	for pos, line := range valueArr {
		if line.name != "Reserved" {
			valName.Add(pos+2, line.name)
			valBegin.Add(pos+2, line.begin)
			valEnd.Add(pos+2, line.end)
			valComment.Add(pos+2, line.comment)
		}
	}

	//

	build(obj, valName, valBegin, valEnd, valComment)
}

func syscon() {
	name := "SysCon"

	bufArr := readCSV(name)

	var valueArr []adrObj
	for pos := 0; pos < len(bufArr); pos++ {
		buf := bufArr[pos]

		obj := adrObj{
			buf[2],
			buf[2],
			buf[0],
			" " + buf[3] + "; " + buf[1],
		}

		if len(bufArr)-1 != pos {
			obj.end = bufArr[pos+1][2]
		} else {
			obj.end = obj.end + "+5"
		}

		valueArr = append(valueArr, obj)
	}

	//

	obj := generator.Init(name, "esp32/address/"+name+".go")

	valName := obj.NewVal()
	valBegin := obj.NewVal()
	valEnd := obj.NewVal()
	valComment := obj.NewVal()

	for pos, line := range valueArr {
		valName.Add(pos+2, line.name)
		valBegin.Add(pos+2, line.begin)
		valEnd.Add(pos+2, line.end)
		valComment.Add(pos+2, line.comment)
	}

	//

	build(obj, valName, valBegin, valEnd, valComment)
}

func DPort() {
	name := "DPort"

	bufArr := readCSV(name)

	var valueArr []adrObj
	for pos := 0; pos < len(bufArr); pos++ {
		buf := bufArr[pos]

		obj := adrObj{
			buf[2],
			buf[2],
			buf[0],
			" " + buf[3] + "; " + buf[1],
		}

		if len(bufArr)-1 != pos {
			obj.end = bufArr[pos+1][2]
		} else {
			obj.end = obj.end + "+32"
		}

		valueArr = append(valueArr, obj)
	}

	//

	obj := generator.Init(name, "esp32/address/"+name+".go")

	valName := obj.NewVal()
	valBegin := obj.NewVal()
	valEnd := obj.NewVal()
	valComment := obj.NewVal()

	for pos, line := range valueArr {
		valName.Add(pos+2, line.name)
		valBegin.Add(pos+2, line.begin)
		valEnd.Add(pos+2, line.end)
		valComment.Add(pos+2, line.comment)
	}

	//

	build(obj, valName, valBegin, valEnd, valComment)
}
