/*
code from www.361way.com
read whole file use ReadFile func
*/

package main

import (
	"fmt"
	"io/ioutil"
	//"strings"
)

func main() {
	Ioutil("x")
}

func Ioutil(name string) {
	if contents, err := ioutil.ReadFile(name); err == nil {
		if len(contents) > 0 {
			s := contents[:len(contents)-1]
			fmt.Println(string(s))
		} else {
			fmt.Println(string(contents))
		}
	}
}
