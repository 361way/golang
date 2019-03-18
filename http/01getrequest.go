/*
  Http (curl) request in golang
  @author www.361way.com <itybku@139.com>
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://361way.com/api/users"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("cache-control", "no-cache")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}


/* 也可以结合http请求和http服务,如下代码：
package main
 
import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
    url := "http://country.io/capital.json"
    response, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()
 
    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
 
    responseString := string(responseData)
    fmt.Fprint(w, responseString)
}
func main() {
    http.HandleFunc("/", ServeHTTP)
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

*/

