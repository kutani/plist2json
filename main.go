/*
 * See LICENSE for licensing information
 */
package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: plist2json [file ...]\n")
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	argv := flag.Args()

	var p *Dict

	if len(argv) == 0 {
		p = ReadPlist(os.Stdin)
		if p == nil {
			fmt.Println("ERROR")
		} else {
			p.Print()
		}
	} else {
		for _, path := range argv {
			f, err := os.Open(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "open %s: %s", path, err)
				continue
			}
			p = ReadPlist(f)
			if p == nil {
				fmt.Println("ERROR")
			} else {
				p.Print()
			}
			f.Close()
		}
	}
}
