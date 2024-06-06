package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func getmyip(addr string) string {
	resp, err := http.Get("http://" + addr + "/myip")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return (string(body))
}

func usage() {
	fmt.Fprintf(os.Stderr, `Get the dcn or work ipaddress,default print all addresses(dcn and work)
Usage: getmyip [-dcn] [-yw]         power by ITO<yangbk> 

Options:
`)
	flag.PrintDefaults()
}

func main() {
	dcn := flag.Bool("dcn", false, "return the host dcn address")
	yw := flag.Bool("yw", false, "return the host work address")

	flag.CommandLine.Usage = func() {
		usage()
		//fmt.PrintDefaults()
	}

	flag.Parse()

	if *dcn {
		fmt.Println(getmyip("10.212.149.204"))
	} else if *yw {
		fmt.Println(getmyip("192.168.23.148"))
	} else {
		fmt.Println("DCN:" + getmyip("10.212.149.204"))
		fmt.Println("YW:" + getmyip("192.168.23.148"))
	}
}
