package main

import (
	"encoding/json"
	"fmt"
	//"io/ioutil"
	//"log"
	"net/http"
	"strings"
	"time"
)

type DomainData struct {
	Action string `json:"action"`
	Node   Node   `json:"node"`
}
type Nodes struct {
	Key           string `json:"key"`
	Value         string `json:"value"`
	ModifiedIndex int    `json:"modifiedIndex"`
	CreatedIndex  int    `json:"createdIndex"`
}
type Node struct {
	Key           string  `json:"key"`
	Dir           bool    `json:"dir"`
	Nodes         []Nodes `json:"nodes"`
	ModifiedIndex int     `json:"modifiedIndex"`
	CreatedIndex  int     `json:"createdIndex"`
}

func reverse(strs []string) []string {
	newStrs := make([]string, len(strs))
	for i, j := 0, len(strs)-1; i <= j; i, j = i+1, j-1 {
		newStrs[i], newStrs[j] = strs[j], strs[i]
	}
	return newStrs
}

func getJson(url string, target interface{}) error {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
func main() {
	//fmt.Println(strings.Split("68.v-40.binjiang.migu", ".")) //[a b c d e]
	s := strings.Split("68.v-40.binjiang.migu", ".") //[a b c d e]
	qurl := strings.Join(reverse(s), "/")
	url := "http://192.168.27.61:2379/v2/keys/skydns/" + qurl
	fmt.Println(url)

	data := DomainData{}
	getJson(url, &data)
	fmt.Println(data)
}
