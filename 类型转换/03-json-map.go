package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := `
        {
            "name": "yang",
            "age": 18
        }
        `
	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &mapResult)
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
	}
	fmt.Println(mapResult)
	fmt.Println(mapResult["name"])
}
