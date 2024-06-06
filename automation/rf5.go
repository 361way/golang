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
	"strings"
)

func NewCFBEncrypter(text string) string {
	key, _ := hex.DecodeString("6368616e676520746869732070610808")
	fmt.Println("key value is:",key)
	plaintext := []byte(text)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	fmt.Println("block value is:",block)
	fmt.Println("aes block size:",aes.BlockSize)
	fmt.Println("plaintext len size:",len(plaintext))

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	fmt.Println("iv value is:",iv)
	fmt.Println("string iv value is:",string(iv))
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
	return (string(contents))
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
		starray := os.Args[2:]
		st := strings.Join(starray," ")
		fmt.Printf("You input sting is:%s\n", st)
    fmt.Print("And command Encrypter output is:")
		c :=Cmd{}
		c.Command = st
		if cmdJSON, err := json.Marshal(c); err == nil {
				fmt.Println(string(cmdJSON)+"\n")
				NewCFBEncrypter(string(cmdJSON)) 
		}
	}
	if os.Args[1] == "-f" {
		st := os.Args[2]
		fmt.Printf("You input sting is:%s\n", Ioutil(st))
		fmt.Print("And file Encrypter output is:")

		c := Cmd{}
		c.Command = Ioutil(st)
		if cmdJSON, err := json.Marshal(c); err == nil {
      
			fmt.Println(string(cmdJSON)+"\n")

			NewCFBEncrypter(string(cmdJSON))
		}
	}
}
