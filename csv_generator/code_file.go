package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/huoshan017/go_libjmy/table"
)

type CodeFile struct {
	reader    *table.CsvReader
	file_name string
	code      string
}

const (
	CSV_HEADER_ROWS_NUM           = 2
	CSV_HEADER_COLUMNS_NAME_INDEX = 0
	CSV_HEADER_COLUMNS_TYPE_INDEX = 1
)

func (this *CodeFile) GetFileName() string {
	return this.file_name
}

func (this *CodeFile) LoadCsv(file string) bool {
	cr := &table.CsvReader{}
	if !cr.Load(file) {
		fmt.Printf("load csv file %v failed\n", file)
		return false
	}

	dot_idx := strings.Index(file, ".csv")
	if dot_idx < 0 {
		fmt.Printf("get index for file name %v failed\n", file)
		return false
	}

	for i := dot_idx; i >= 0; i-- {
		a := []byte(file)[i]
		if a == byte('/') || a == byte('\\') {
			this.file_name = string([]byte(file)[i+1 : dot_idx])
			break
		}
	}

	if this.file_name == "" {
		fmt.Printf("cant get file name from string[%v]\n", file)
		return false
	}

	this.reader = cr
	fmt.Printf("load csv file %v success\n", file)
	return true
}

func (this *CodeFile) Generate() bool {
	if this.reader == nil {
		return false
	}

	names_row := this.reader.GetRow(CSV_HEADER_COLUMNS_NAME_INDEX)
	if names_row == nil {
		fmt.Printf("csv file %v name row not found\n", this.file_name)
		return false
	}

	types_row := this.reader.GetRow(CSV_HEADER_COLUMNS_TYPE_INDEX)
	if types_row == nil {
		fmt.Printf("csv file %v type row not found\n", this.file_name)
		return false
	}

	if names_row.GetItemsNum() == 0 {
		fmt.Printf("csv file %v not include items\n", this.file_name)
		return false
	}

	if names_row.GetItemsNum() != types_row.GetItemsNum() {
		fmt.Printf("csv file %v name row length[%v] not same to type row length[%v]\n", this.file_name, names_row.GetItemsNum(), types_row.GetItemsNum())
		return false
	}

	row_num := strconv.Itoa(CSV_HEADER_ROWS_NUM)
	names_index := strconv.Itoa(CSV_HEADER_COLUMNS_NAME_INDEX)
	types_index := strconv.Itoa(CSV_HEADER_COLUMNS_NAME_INDEX)

	this.code = "package main\n\n"
	this.code += "import(\n"
	this.code += "  \"fmt\"\n"
	this.code += "  \"strconv\"\n"
	this.code += "  \"github.com/huoshan017/go_libjmy/table\"\n"
	this.code += ")\n\n"

	// csv row struct
	row_struct_name := this.file_name + "Item"
	this.code += ("type " + row_struct_name + " struct {\n")
	for i := 0; i < names_row.GetItemsNum(); i++ {
		this.code += ("  " + names_row.GetItem(i) + " " + types_row.GetItem(i) + "\n")
	}
	this.code += "}\n\n"

	// csv struct
	this.code += ("type " + this.file_name + " struct {\n")
	// 0 index item is the default key
	this.code += ("  map_items map[" + types_row.GetItem(0) + "]*" + row_struct_name + "\n")
	// array items
	this.code += ("  arr_items []*" + row_struct_name + "\n")
	this.code += "}\n\n"

	// load function
	this.code += ("func (this *" + this.file_name + ") Load(file string) bool {\n")
	this.code += ("  cr := &table.CsvReader{}\n")
	this.code += ("  if !cr.Load(file) {\n")
	this.code += ("    fmt.Println(\"load csv file %v failed\", file)\n")
	this.code += ("    return false\n")
	this.code += ("  }\n\n")
	this.code += ("  data_row_num := cr.GetRowNum() - " + row_num + "\n")
	this.code += ("  if data_row_num <= 0 {\n")
	this.code += ("    fmt.Println(\"no data row\")\n")
	this.code += ("    return false\n")
	this.code += ("  }\n\n")
	this.code += ("  this.map_items = make(map[" + types_row.GetItem(0) + "]*" + row_struct_name + ", data_row_num)\n")
	this.code += ("  this.arr_items = make([]*" + row_struct_name + ", data_row_num)\n")
	this.code += ("  names_row := cr.GetRow(" + names_index + ")\n")
	this.code += ("  if names_row == nil {\n")
	this.code += ("    fmt.Println(\"names_row is null\")\n")
	this.code += ("    return false\n")
	this.code += ("  }\n")
	this.code += ("  types_row := cr.GetRow(" + types_index + ")\n")
	this.code += ("  if types_row == nil {\n")
	this.code += ("    fmt.Println(\" types_row is null\")\n")
	this.code += ("    return false\n")
	this.code += ("  }\n\n")
	this.code += ("  for i:=0; i<data_row_num; i++ {\n")
	this.code += ("    r := &" + row_struct_name + "{}\n")
	this.code += ("    data_row := cr.GetRow(i + " + strconv.Itoa(CSV_HEADER_ROWS_NUM) + ")\n")
	this.code += ("    v := 0\n")
	this.code += ("    var e error\n")
	this.code += ("    e = nil\n")
	for j := 0; j < names_row.GetItemsNum(); j++ {
		row_item := "data_row.GetItem(" + strconv.Itoa(j) + ")"
		if types_row.GetItem(j) == "int64" {
			this.code += ("    v, e = strconv.ParseInt(" + row_item + ", 10, 64)\n")
			this.code += ("    if e != nil {\n")
			this.code += ("      fmt.Printf(\"parse int64 value failed\\n\")\n")
			this.code += ("      return false\n")
			this.code += ("    }\n")
			this.code += ("    r." + names_row.GetItem(j) + " = strconv.FormatInt(e, 10)\n")
		} else if types_row.GetItem(j) == "string" {
			this.code += ("    r." + names_row.GetItem(j) + " = " + row_item + "\n")
		} else {
			this.code += ("    v, e = strconv.Atoi(" + row_item + ")\n")
			this.code += ("    if e != nil {\n")
			this.code += ("      fmt.Printf(\"string to int failed\\n\")\n")
			this.code += ("      return false\n")
			this.code += ("    }\n")
			this.code += ("    r." + names_row.GetItem(j) + " = " + types_row.GetItem(j) + "(v)\n")
		}
	}
	this.code += ("    this.map_items[r." + names_row.GetItem(0) + "] = r\n")
	this.code += ("    this.arr_items[i] = r\n")
	this.code += ("  }\n")
	this.code += ("  return true")
	this.code += ("}\n\n")

	// close function
	this.code += ("func (this *" + this.file_name + ") Close() {\n")
	this.code += "}\n\n"

	// get row item function by key
	this.code += ("func (this *" + this.file_name + ") Get(key " + types_row.GetItem(0) + ") *" + row_struct_name + " {\n")
	this.code += ("  v, o := this.map_items[key]\n")
	this.code += ("  if !o {\n")
	this.code += ("    return nil\n")
	this.code += ("  }\n")
	this.code += ("  return v\n")
	this.code += ("}\n\n")

	// get row item function by index
	this.code += ("func (this *" + this.file_name + ") GetByIndex(index int) *" + row_struct_name + " {\n")
	this.code += ("  if index >= len(this.arr_items) {\n")
	this.code += ("    return nil\n")
	this.code += ("  }\n")
	this.code += ("  return this.arr_items[index]\n")
	this.code += ("}")

	return true
}

func (this *CodeFile) Save(dest_file string) bool {
	if this.reader == nil {
		fmt.Println("not read csv file yet")
		return false
	}
	f, e := os.Create(dest_file)
	if e != nil {
		fmt.Printf("create dest file %v failed, err %v\n", dest_file, e.Error())
		return false
	}

	_, e = f.WriteString(this.code)
	if e != nil {
		fmt.Printf("write string for dest file %v failed\n", dest_file)
		return false
	}

	return true
}
