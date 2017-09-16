package table

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type CsvRow struct {
	items []string
}

func (this *CsvRow) GetItemsNum() int {
	return len(this.items)
}

func (this *CsvRow) GetItem(index int) string {
	if index < 0 || index >= len(this.items) {
		return ""
	}
	return this.items[index]
}

type CsvReader struct {
	rows []*CsvRow
}

func (this *CsvReader) Load(file string) bool {
	fi, err := os.Open(file)
	if err != nil {
		fmt.Println("load csv file %v failed, err: %v", file, err.Error())
		return false
	}
	defer fi.Close()

	br := bufio.NewReader(fi)

	for {
		a, _, e := br.ReadLine()
		if e == io.EOF {
			break
		}

		items := strings.Split(string(a), ",")
		if items == nil || len(items) == 0 {
			fmt.Println("split content is empty")
			return false
		}

		row := &CsvRow{}
		row.items = items

		if this.rows == nil {
			this.rows = make([]*CsvRow, 0)
		}

		this.rows = append(this.rows, row)
	}

	return true
}

func (this *CsvReader) Close() {
	if this.rows == nil {
		return
	}

	for i, _ := range this.rows {
		this.rows[i] = nil
	}

	this.rows = nil
}

func (this *CsvReader) GetRowNum() int {
	if this.rows == nil {
		return 0
	}
	return len(this.rows)
}

func (this *CsvReader) GetRow(index int) *CsvRow {
	if this.rows == nil || index >= len(this.rows) {
		return nil
	}
	return this.rows[index]
}
