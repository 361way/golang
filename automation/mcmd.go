package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

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
	c := Cmd{}

	n1 := len(*command) > 0
	n2 := len(*ip) > 0
	n3 := len(*filename) > 0
	//n4 := n2 && n1
	//n5 := n2 && n3
	n6 := len(*ipfile) > 0
	//n7 := n1 && n6
	//n8 := n3 && n6

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
		//fmt.Println(string(cmdJSON))
		//data := []byte("abacba2d2f2e4b01c8ab058f147cb61caa05acf5401f7e47d725e5629bbbbd12e35f9478")
		//data := []byte(ExampleNewCFBEncrypter(st))
		if n2 {
			//url := "http://10.212.52.252:15380"
			url := "http://" + *ip + *port
			x, _ := DoBytesPost(url, data)
			fmt.Println(string(x))
		} else if n6 {
			var wg sync.WaitGroup
			lines := strings.Split(Ioutil(*ipfile), "\n")
			//fmt.Println(lines)
			for _, r := range lines {
				r = strings.TrimSpace(r)
				if len(r) > 6 {
					wg.Add(1)
					//fmt.Println("value is: " + r)
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

func Multrun(ip string, data []byte, wg *sync.WaitGroup) {
	defer wg.Done()
	url := "http://" + ip + ":15380"
	x, _ := DoBytesPost(url, data)
	fmt.Println(ip + "\n================================\n" + string(x))
}

//func main() {
//	done := make(chan bool)
//	done2 := make(chan bool)
//	go Readip(done, ipch, *ipfile)
//	<-done
//}
