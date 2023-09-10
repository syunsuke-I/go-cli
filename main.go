package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	showAll := flag.Bool("a", false, "all files")
	showDetail := flag.Bool("l", false, "show detail info")
	flag.Parse()
	items, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		file_name := item.Name()
		is_hidden_file := strings.HasPrefix(file_name, ".")
		switch {
		case *showDetail:
			file_info, err := item.Info()
			if err != nil {
				log.Fatal(err)
			}
			file_size := file_info.Size()
			mod_time := file_info.ModTime()
			if !is_hidden_file {
				// fmt.Printf("%10s\t%d\t%s\n", file_name, file_size, mod_time)
				fmt.Printf("%s\t%10d%10s\n", mod_time, file_size, file_name)
			}
			// 隠しファイルは表示しない
		case *showAll:
			println(file_name)
		default:
			if !is_hidden_file {
				println(file_name)
			}
		}
	}
}
