package main

//引入相关模块
import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/donnie4w/go-logger/logger"
)

//定义结构体User和Parameter，并指定相关参数类型

type User struct {
	Command string
}

type Parameter struct {
	//	Log      bool
	Path     string
	FileName string
}

const ShellToUse = "bash"

//定义命令执行函数Shellout，该函数会将执行的结果返回
//其返回有执行错误、标准输出、标准错误
func Shellout(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}

func ExecFile(program string, filetype string) string {

	var result string
	fileType := "./extend/" + filetype
	matches, _ := filepath.Glob(fileType)
	for _, v := range matches {
		var stdout, stderr bytes.Buffer

		cmd := exec.Command(program, v)
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			//log.Fatal(v, " script exec error is:", err)
			log.Print(v, " script exec error is:", err)
		}
		if len(stderr.String()) > 0 {
			log.Print(v, " script exec std error is:", stderr.String())
		}
		result = result + "# script " + v + " exec result is:\n" + stdout.String()
	}

	return result
}

func Sysinfo() string {
	cmd := `
	echo "## date"
	date +'%Y-%m-%d %T'
	echo "## meminfo"
	free -m
	echo "## cpuinfo"
	top -b -n1 | grep -B1 '%Cpu'
	echo "## uptime-load"
	uptime
	echo "## disk"
	df -hP
	echo "## inode"
	df -hPi
	`
	_, sysinfo, _ := Shellout(cmd)
	return sysinfo

}

func TcpSconn(Sdata string) {

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("dial failed:", err)
		//os.Exit(1)
	} else {
		defer conn.Close()
		// send to socket
		//text := []byte("111111111111111111111111111\n2222xyzdrwrewfdisn")
		//sysdata := Sysinfo()
		text := []byte(Sdata)
		encodedStr := hex.EncodeToString(text)
		//fmt.Println(encodedStr)
		fmt.Fprintf(conn, encodedStr+"\n")
	}

}

func SendData() {
	tc := time.Tick(30 * time.Second)
	for {
		data := ExecFile("sh", "*.sh")
		fmt.Println(data)
		data2 := ExecFile("python", "*.py")
		fmt.Println(data2)
		sysdata := Sysinfo()
		d := data + data2 + sysdata
		TcpSconn(d)
		<-tc
	}

}

//定义解密函数，控制端加密后的结果会通过该函数解密并输出相关明文
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

//定义一个无日志的http监听函数(修复handler函数传入/dev/null异常的问题)
func nolog(w http.ResponseWriter, r *http.Request) {
	var u User
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	body, _ := ioutil.ReadAll(r.Body)
	d := NewCFBDecrypter(string(body))
	b := bytes.NewBuffer(d)
	err := json.NewDecoder(b).Decode(&u)
	//fmt.Println(u)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	Cip := strings.Split(r.RemoteAddr, ":")
	//fmt.Println(Cip[0])
	if Cip[0] == "200.200.30.103" || Cip[0] == "200.200.30.104" {
		err, out, errout := Shellout(u.Command)
		if err != nil {
			fmt.Println("error: %v\n", err)
		}
		w.Write([]byte(out))
		fmt.Println(errout)
	}
}

func Collect(w http.ResponseWriter, r *http.Request) {
	data := ExecFile("sh", "*.sh")
	data2 := ExecFile("python", "*.py")
	sysdata := Sysinfo()
	d := data + data2 + sysdata
	w.Write([]byte(d))
}

//定义一个http监听函数，后面会调用到
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

	//	确认是否在终端界面打印日志：
	//	true代表打印
	//	false代表不在终端打印
	logger.SetConsole(false)
	//	var lg = logger.GetLogger()
	//	lg.SetLevel(logger.INFO)
	//	lg.SetLevelFile(logger.INFO, `/tmp/`, "info.log")
	//	lg.SetLevelFile(logger.WARN, `/tmp/`, "warn.log")

	//	定义日志输出的级别和轮询输出的路径、文件名
	//	如果想要定义多个日志文件，而且存放到不同路径，可以参考上面注释的用法
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

	//	来源主机控制，非允许主机，不允许连接并执行命令
	//	相比于UUID、session、动态码等方式而言，更加简单高效
	//	从测试数据来看，效率较前几种确实高，而且控制方式更加细度
	if Cip[0] == "200.200.30.103" || Cip[0] == "200.200.30.104" {
		err, out, errout := Shellout(u.Command)
		if err != nil {
			//fmt.Println("error: %v\n", err)
			logger.Error(err)
		}
		w.Write([]byte(out))
		//fmt.Println(errout)
		if len(errout) > 1 {
			logger.Error(errout)
		}
	}
}

//主函数
func main() {

	//	参数变量及其默认值
	logs := flag.Bool("logs", false, "a bool value: true or false")
	port := flag.String("port", ":10000", "run port")
	flag.Parse()
	if *logs {
		myHandler := &Parameter{Path: "./", FileName: "info.log"}
		http.HandleFunc("/", myHandler.handler)
	} else {
		http.HandleFunc("/", nolog)
	}
	http.HandleFunc("/collect", Collect)
	//	http.HandleFunc("/", myHandler.handler)
	err := http.ListenAndServe(*port, nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
	SendData()

}
