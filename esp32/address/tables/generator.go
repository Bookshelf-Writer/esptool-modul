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

func main() {
	peripheral()
}

////

func peripheral() {
	name := "peripheral"

	bufArr := readCSV(name)
	type PeripheralObj struct {
		begin   string
		end     string
		name    string
		comment string
	}

	var valueArr []PeripheralObj
	for _, buf := range bufArr {
		if len(buf[0]) > 0 {
			obj := PeripheralObj{
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

	obj := generator.Init(name[:3], "esp32/address/"+name+".go")

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

	obj.PrintLN("const (")
	for _, code := range valName.Get.Ints() {
		text := valName.Get.Text(code)

		obj.Offset(1).Name.SelfCode(text).Print(" ")
		obj.Name.Type().Print(" = ").Number(code)
		obj.Offset(1).Print("//" + valComment.Get.Text(code))
		obj.LN()
	}
	obj.PrintLN(")").LN()

	//

	obj.PrintLN("const (")
	for _, code := range valName.Get.Ints() {
		text := valName.Get.Text(code)

		obj.Offset(1).Name.SelfParam("AdrBegin", text)
		obj.Print(" AddressType = ")
		obj.Print(valBegin.Get.Text(code)).LN()
	}
	obj.PrintLN(")").LN()

	obj.PrintLN("const (")
	for _, code := range valName.Get.Ints() {
		text := valName.Get.Text(code)

		obj.Offset(1).Name.SelfParam("AdrEnd", text)
		obj.Print(" AddressType = ")
		obj.Print(valEnd.Get.Text(code)).LN()
	}
	obj.PrintLN(")").LN()

	obj.PrintLN("var (")
	for _, code := range valName.Get.Ints() {
		text := valName.Get.Text(code)

		obj.Offset(1).Name.SelfParam("Adr", text)
		obj.Print(" = AddressCellObj{")
		obj.Name.SelfParam("AdrBegin", text).Print(",")
		obj.Name.SelfParam("AdrEnd", text).PrintLN("}")
	}
	obj.PrintLN(")").LN()

	//

	obj.Print("var ").Name.Map().Print(" = map[").Name.Type().PrintLN("]*AddressCellObj{")
	for _, code := range valName.Get.Ints() {
		text := valName.Get.Text(code)

		obj.Offset(1).Name.SelfCode(text).Print(": &")
		obj.Name.SelfParam("Adr", text).PrintLN(",")

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
