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
		fmt.Println("load csv file %v failed", file)
		return false
	}

	dot_idx := strings.Index(file, ".csv")
	if dot_idx < 0 {
		fmt.Println("get index for file name %v failed", file)
		return false
	}

	for i := dot_idx; i >= 0; i-- {
		a := []byte(file)[i]
		if a == byte('/') {
			this.file_name = string([]byte(file)[i+1 : dot_idx])
			break
		}
	}

	if this.file_name == "" {
		fmt.Println("cant get file name from string[%v]", file)
		return false
	}

	this.reader = cr
	fmt.Println("load csv file %v success", file)
	return true
}

func (this *CodeFile) Generate() bool {
	if this.reader == nil {
		return false
	}

	names_row := this.reader.GetRow(CSV_HEADER_COLUMNS_NAME_INDEX)
	if names_row == nil {
		fmt.Println("csv file %v name row not found", this.file_name)
		return false
	}

	types_row := this.reader.GetRow(CSV_HEADER_COLUMNS_TYPE_INDEX)
	if types_row == nil {
		fmt.Println("csv file %v type row not found", this.file_name)
		return false
	}

	if names_row.GetItemsNum() == 0 {
		fmt.Println("csv file %v not include items", this.file_name)
		return false
	}

	if names_row.GetItemsNum() != types_row.GetItemsNum() {
		fmt.Println("csv file %v name row length[%v] not same to type row length[%v]", this.file_name, names_row.GetItemsNum(), types_row.GetItemsNum())
		return false
	}

	num := this.reader.GetRowNum()
	this.code = "package main\n\n"
	this.code += "import(\n"
	this.code += "  \"fmt\"\n"
	this.code += "  \"io\"\n"
	this.code += "  \"os\"\n"
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
	this.code += ("  arr_items []*" + row_struct_name + "\n")
	this.code += "}\n\n"

	// load function
	this.code += ("func (this *" + this.file_name + ") Load(file string) bool {\n")
	this.code += ("  cr := &table.CsvReader{}\n")
	this.code += ("  if !cr.Load(file) {\n")
	this.code += ("    fmt.Println(\"load csv file %v failed\", file)\n")
	this.code += ("    return false\n")
	this.code += ("  }\n\n")
	this.code += ("  data_row_num := cr.GetRowNum() - " + strconv.Itoa(CSV_HEADER_ROWS_NUM) + "\n")
	this.code += ("  if data_row_num <= 0 {\n")
	this.code += ("    fmt.Println(\"no data row\")\n")
	this.code += ("    return false\n")
	this.code += ("  }\n\n")
	this.code += ("  ")
	this.code += ("  return true")
	this.code += ("}\n\n")

	// close function
	this.code += ("func (this *" + this.file_name + ") Close() {\n")
	this.code += "}\n\n"

	return true
}

func (this *CodeFile) Save(dest_file string) bool {
	if this.reader == nil {
		fmt.Println("not read csv file yet")
		return false
	}
	f, e := os.Create(dest_file)
	if e != nil {
		fmt.Println("create dest file %v failed, err %v", dest_file, e.Error())
		return false
	}

	_, e = f.WriteString(this.code)
	if e != nil {
		fmt.Println("write string for dest file %v failed", dest_file)
		return false
	}

	return true
}
