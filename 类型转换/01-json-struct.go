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
	jsonStr := `{
            "name_title": "yang",
            "age_size":30
        }`
	p := People{}
	err := json.Unmarshal([]byte(jsonStr), &p)
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
	fmt.Println(p.Age)
}
