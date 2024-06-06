package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Name string `json:"name_title"`
	Age  int    `json:"age_size"`
}

func main() {
	p := People{
		Name: "yang",
		Age:  18,
	}
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonBytes))
}
