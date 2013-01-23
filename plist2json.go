/*
 * See LICENSE for licensing information
 */
package main

import (
	"fmt"
	"flag"
	"os"
	"bufio"
)

func main() {
	flag.Parse()
	argv := flag.Args()
	argc := len(argv)

	if argc < 1 {
		fmt.Println("No file passed")
		return
	}

	f, err := os.Open(argv[0])
	defer f.Close()

	if err != nil {
		fmt.Println("Error loading file: ",err)
		return
	}

	p := ReadPlist(bufio.NewReader(f))

	if p == nil {
		fmt.Println("ERROR")
	}

	p.Print()
}




