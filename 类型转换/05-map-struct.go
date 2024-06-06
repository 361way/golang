package main

import (
	"fmt"

	"github.com/goinggo/mapstructure"
)

type People struct {
	Name string `json:"name_title"`
	Age  int    `json:"age_size"`
}

func main() {
	mapInstance := make(map[string]interface{})
	mapInstance["Name"] = "jqw"
	mapInstance["Age"] = 18

	var people People
	err := mapstructure.Decode(mapInstance, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)
}
