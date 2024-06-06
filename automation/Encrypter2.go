package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	//"strings"
)

func NewCFBEncrypter(text string) string {
	key, _ := hex.DecodeString("6368616e676520746869732070610808")
	plaintext := []byte(text)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	fmt.Printf("%x\n", ciphertext)
	return (string(ciphertext))
}

func Ioutil(name string) string {
	contents, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
	//	result := strings.Replace(string(contents), "\n", "", 1)
	//	fmt.Println(result)
	//	return (result)
	fmt.Println(contents)
	return (contents)
}

type Cmd struct {
	Command string
}

func main() {

	if len(os.Args) != 3 {
		fmt.Print("usage:\n")
		fmt.Printf("\t -c command\n")
		fmt.Printf("\t -f filename\n")
		fmt.Printf("example:\n")
		fmt.Printf("\t%s -c '{\"Command\": \"df -h\"}' \n", os.Args[0])
		fmt.Printf("\t%s -f test.sh\n", os.Args[0])
		os.Exit(1)
	}

	if os.Args[1] == "-c" {
		st := os.Args[2]
		fmt.Printf("You input sting is:%s\n", st)
		fmt.Print("And Encrypter output is:")
		NewCFBEncrypter(st)
	}
	if os.Args[1] == "-f" {
		st := os.Args[2]
		fmt.Printf("You input sting is:%s\n", Ioutil(st))
		fmt.Print("And Encrypter output is:")

		c := Cmd{}
		c.Command = Ioutil(st)
		if cmdJSON, err := json.Marshal(c); err == nil {

			NewCFBEncrypter(string(cmdJSON))
			fmt.Println(string(cmdJSON))
		}
	}
}
