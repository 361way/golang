package main 
    import (
        "net/http"
        "encoding/json"
        "bytes"
        "fmt"
    )
type User struct{
    Id      string
    Balance uint64
}

func main() {
        u := User{Id: "www.361way.com", Balance: 8}
        b := new(bytes.Buffer)
        json.NewEncoder(b).Encode(u)
 			  res, _ := http.Post("https://httpbin.org/post", "application/json; charset=utf-8", b)
        var body struct {
                //sends back key/value pairs, no map[string][]string
                Headers map[string]string `json:"headers"`
                Origin  string            `json:"origin"`
        }
        json.NewDecoder(res.Body).Decode(&body)
        fmt.Println(body)
}
