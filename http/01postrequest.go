/*
  Http (curl) request in golang
  @author www.361way.com <itybku@139.com>
*/
package main

import (
"fmt"
"net/http"
"io/ioutil"
)

func main() {
	url := "https://reqres.in/api/users"
	payload := strings.NewReader("name=test&jab=teacher")
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cache-control", "no-cache")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
