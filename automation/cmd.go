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
        ip := flag.String("ip", "", "the will be exec host ip address")
        port := flag.String("port", ":15380", "the will be exec host ip port")
        flag.Parse()
        c := Cmd{}
        //url := "http://10.212.52.252:15380"
        url := "http://" + *ip + *port

        n1 := len(*command) > 0
        n2 := len(*ip) > 0
        n3 := len(*filename) > 0
        n4 := n2 && n1
        n5 := n2 && n3

        if n4 {
                c.Command = *command
        } else if n5 {
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
                x, _ := DoBytesPost(url, data)
                fmt.Println(string(x))
        }

}

