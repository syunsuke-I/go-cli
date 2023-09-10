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
	flag.Parse()
	items, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range items {
		file_name := item.Name()
		// 隠しファイルは表示しない
		if strings.HasPrefix(file_name, ".") {
			if *showAll {
				println(file_name)
				continue
			}
			continue
		}
		fmt.Println(file_name)
	}
}
