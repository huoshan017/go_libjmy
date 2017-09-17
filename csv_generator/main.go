package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dest_path := flag.String("d", "./", "dest file path")
	flag.Parse()
	if dest_path == nil {
		default_dest := "./"
		dest_path = &default_dest
	}

	fmt.Println("dest path ", *dest_path)

	csv_path := "../conf/csv_generator/csv"
	filepath.Walk(csv_path, func(path string, f os.FileInfo, err error) error {
		fmt.Printf("path is %v\n", path)
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		if strings.Index(f.Name(), ".csv") < 0 {
			return nil
		}

		fmt.Printf("file name is %v\n", path)

		cf := &CodeFile{}
		if !cf.LoadCsv(path) {
			err_str := "invalid csv file " + path
			return errors.New(err_str)
		}
		if !cf.Generate() {
			err_str := "generate code failed with csv file " + path
			return errors.New(err_str)
		}

		file_name := *dest_path + "\\" + cf.GetFileName() + ".go"
		if !cf.Save(file_name) {
			err_str := "save code file " + path + " to dest failed"
			return errors.New(err_str)
		}
		fmt.Printf("saved new code file %v\n", file_name)
		return nil
	})
}
