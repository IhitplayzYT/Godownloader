package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	l := len(args)
	if l < 3 || l&1 == 0 {
		usage(args[0])
		return
	}
	url_path_map := parse_args(args[1:])
	fmt.Println(url_path_map)
	for url, path := range url_path_map {
		err := download(url, path)
		if err != nil {
			fmt.Println("Error Occured", err)
		}
	}

}
