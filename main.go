package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	items, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range items {
		// ファイルの情報を取得
		file_info, err := item.Info()
		if err != nil {
			log.Fatal(err)
		}

		file_name := item.Name()
		file_size := file_info.Size()
		file_name_init := file_name[0:1]
		// 隠しファイルは表示しない
		if file_name_init != "." {
			fmt.Printf("%s\t%d\n", file_name, file_size)
		}
	}

}
