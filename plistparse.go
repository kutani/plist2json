/*
 * See LICENSE for licensing information
 */
package main

import (
	"fmt"
	"bufio"
	"strconv"
)

type Dict struct {
	array []*KeyPair
}

func (self *Dict) add(k *KeyPair) {
	l := len(self.array)

	n_arry := make([]*KeyPair, l+1)
	copy(n_arry, self.array)
	self.array = n_arry

	self.array = self.array[0:l+1]
	self.array[l] = k
}

func (self *Dict) Print() {
	fmt.Print("{")
	l := len(self.array)
	for i := 0; i < l; i++ {
		if i > 0 && i != l {
			fmt.Print(",")
		}
		self.array[i].Print()
	}
	fmt.Print("}")
}

type KeyPair struct {
	k string
	t string
	s string
	b string
	i int
	d *Dict
}

func (self *KeyPair) Print() {
	fmt.Print("\""+self.k+"\": ")
	if self.t == "string" {
		fmt.Print("\""+self.s+"\"")
		return
	}

	if self.t == "bool" {
		fmt.Print(self.b)
		return
	}

	if self.t == "integer" {
		fmt.Print(self.i)
	}

	if self.t == "dict" {
		self.d.Print()
	}
}


/*
func ReadPlist(rd *bufio.Reader) *Dict {

	var d *Dict = nil

	for  {
		c, s, err := rd.ReadRune()

		if err != nil {
			fmt.Println(err)
			break
		}

		if s > 1 {
			fmt.Println("Got Rune, but we don't support unicode yet!")
			break
		}


		switch(c) {
		case '<':
			st := rd.getNextToken('>')
			fmt.Println(st)

			if st == "dict" {
				if d == nil {
					d = new(Dict)
					continue
				}
			} else if st == "key" {
				k := new(KeyPair)
				st := rd.getNextToken('<')
				k.k = st
			}
		}
	}


	return d
}*/

func ReadPlist(rd *bufio.Reader) *Dict {
	for {
		c, s, err := rd.ReadRune()

		if err != nil {
			fmt.Println(err)
			break
		}

		if s > 1 {
			fmt.Println("Got Rune, but we don't support unicode yet!")
			break
		}

		if c == '<' {
			st := getNextToken(rd, '>')
			if st == "/plist" {
				break
			}

			if st == "dict" {
				return readDict(rd)
			}
		}
	}
	return nil
}


func readDict(rd *bufio.Reader) *Dict {

	d := new(Dict)
	var k *KeyPair

	for {
		c, s, err := rd.ReadRune()
		if err != nil {
			fmt.Println(err)
			break
		}
		if s > 1 {
			fmt.Println("Got Rune, but we don't support unicode yet!")
			break
		}

		if c == '<' {
			st := getNextToken(rd, '>')
			if st == "/dict" {
				break
			}

			if st == "key" {
				k = new(KeyPair)
				d.add(k)
				st := getNextToken(rd,'<')
				k.k = st
				continue
			}

			if st == "dict" {
				k.t = st
				k.d = readDict(rd)
				continue
			}

			if st == "string" {
				k.t = st
				st := getNextToken(rd,'<')
				k.s = st
			}

			if st == "integer" {
				k.t = st
				st := getNextToken(rd,'<')
				is, err := strconv.Atoi(st)
				if err == nil {
					k.i = is 
				} else {
					k.i = 0
				}
			}

			if st == "true/" {
				k.t = "bool"
				k.b = "true"
			}

			if st == "false/" {
				k.t = "bool"
				k.b = "false"
			}

		}

	}
	return d
}


func getNextToken(rd *bufio.Reader, delim byte) string {
	s, err := rd.ReadString(delim)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	s = s[:len(s)-1]
	return s
}


