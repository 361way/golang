package main

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base32"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func toBytes(value int64) []byte {
	var result []byte
	mask := int64(0xFF)
	shifts := [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		result = append(result, byte((value>>shift)&mask))
	}
	return result
}

func toUint32(bytes []byte) uint32 {
	return (uint32(bytes[0]) << 24) + (uint32(bytes[1]) << 16) +
		(uint32(bytes[2]) << 8) + uint32(bytes[3])
}

//func oneTimePassword(key []byte, value []byte) uint32 {
func oneTimePassword(key []byte, value []byte) string {
	hmacSha1 := hmac.New(sha1.New, key)
	hmacSha1.Write(value)
	hash := hmacSha1.Sum(nil)

	offset := hash[len(hash)-1] & 0x0F
	hashParts := hash[offset : offset+4]
	hashParts[0] = hashParts[0] & 0x7F
	number := toUint32(hashParts)
	pwd := number % 1000000

	//return pwd
	return strconv.Itoa(int(pwd))
}

func DoBytesPost(url string, data []byte) ([]byte, error) {

	body := bytes.NewReader(data)
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Println("http.NewRequest,[err=%s][url=%s]", err, url)
		return []byte(""), err
	}
	request.Header.Set("Connection", "Keep-Alive")
	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		log.Println("http.Do failed,[err=%s][url=%s]", err, url)
		return []byte(""), err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("http.Do failed,[err=%s][url=%s]", err, url)
	}
	return b, err
}
func ExampleNewCFBEncrypter(text string) string {
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
	//fmt.Printf("%x\n", ciphertext)
	encodedStr := hex.EncodeToString(ciphertext)
	return (encodedStr)
}

func Ioutil(name string) string {
	contents, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
	return (string(contents))
}

func Multrun(ip string, data []byte, wg *sync.WaitGroup) {
	defer wg.Done()
	url := "http://" + ip + ":15380"
	x, _ := DoBytesPost(url, data)
	fmt.Println(ip + "\n================================\n" + string(x))
}

type Cmd struct {
	Command string
}

func main() {

	command := flag.String("c", "", "the will be exec `command`")
	filename := flag.String("f", "", "the will be exec `filename`")
	ipfile := flag.String("i", "", "hosts `ipfile` name")
	ip := flag.String("ip", "", "the will be exec host ip address")
	port := flag.String("port", ":15380", "the will be exec host ip port")
	flag.Parse()

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter your key:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		os.Exit(1)
	}
	inputNoSpaces := strings.Replace(input, " ", "", -1)
	inputNoSpaces = strings.Replace(inputNoSpaces, "\n", "", -1)

	key, err := base32.StdEncoding.DecodeString("I77D5BE6LHGVRSAV")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	epochSeconds := time.Now().Unix()
	pwd := oneTimePassword(key, toBytes(epochSeconds/30))
	if inputNoSpaces != pwd {
		fmt.Println("You input one Time Password is Wrong!!!")
		os.Exit(1)
	}

	c := Cmd{}

	n1 := len(*command) > 0
	n2 := len(*ip) > 0
	n3 := len(*filename) > 0
	n6 := len(*ipfile) > 0

	if n1 {
		c.Command = *command
	} else if n3 {
		c.Command = Ioutil(*filename)
	} else {
		fmt.Printf("use correct args,use \t%s -h\n check", os.Args[0])
		os.Exit(0)
	}
	//c.Command = Ioutil(st)
	if cmdJSON, err := json.Marshal(c); err == nil {

		data := []byte(ExampleNewCFBEncrypter(string(cmdJSON)))
		if n2 {
			url := "http://" + *ip + *port
			x, _ := DoBytesPost(url, data)
			fmt.Println(string(x))
		} else if n6 {
			var wg sync.WaitGroup
			lines := strings.Split(Ioutil(*ipfile), "\n")
			for _, r := range lines {
				r = strings.TrimSpace(r)
				if len(r) > 6 {
					wg.Add(1)
					go Multrun(r, data, &wg)
				}
			}
			wg.Wait()
		} else {
			fmt.Printf("didn't hava ip address,use \t%s -h\n check", os.Args[0])
			os.Exit(0)
		}
	}
}
