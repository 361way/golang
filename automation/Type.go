package main

import (
	"fmt"
	//"path"
	"bytes"
	"log"
	"os/exec"
	"path/filepath"
)

func ExecFile(program string, filetype string) string {

	var result string
	fileType := "./extend/" + filetype
	matches, _ := filepath.Glob(fileType)
	for _, v := range matches {
		var stdout, stderr bytes.Buffer

		//fmt.Println(program,v)
		cmd := exec.Command(program, v)
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			log.Fatal(v, " script exec error is:", err)
		}
		if len(stderr.String()) >0 {
			log.Fatal(v, " script exec std error is:", stderr.String())
		}
		result = result + "# script " + v + " exec result is:\n" +  stdout.String()
	}

	return result
}

func main() {
	
	data := ExecFile("sh","*.sh")
	fmt.Println(data)
	data2 := ExecFile("python","*.py")
	fmt.Println(data2)

}

