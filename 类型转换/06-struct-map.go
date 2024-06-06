package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name_title"`
	Age  int    `json:"age_size"`
}

func StructToMapDemo(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}
func main() {
	student := Student{10, "jqw", 18}
	data := StructToMapDemo(student)
	fmt.Println(data)
}
