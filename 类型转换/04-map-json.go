package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	mapInstances := []map[string]interface{}{}
	instance_1 := map[string]interface{}{"name": "John", "age": 10}
	instance_2 := map[string]interface{}{"name": "Alex", "age": 12}
	mapInstances = append(mapInstances, instance_1, instance_2)

	jsonStr, err := json.Marshal(mapInstances)

	if err != nil {
		fmt.Println("MapToJsonDemo err: ", err)
	}
	fmt.Println(string(jsonStr))
}
