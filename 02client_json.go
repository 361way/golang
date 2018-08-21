package main 
    import (
        "net/http"
        "encoding/json"
        "io"
        "os"
        "bytes"
    )
type User struct{
    Id      string
    Balance uint64
}
func main() {
        u := User{Id: "www.361way.com", Balance: 8}
        b := new(bytes.Buffer)
        json.NewEncoder(b).Encode(u)
        res, _ := http.Post("http://127.0.0.1:8080", "application/json; charset=utf-8", b)
        io.Copy(os.Stdout, res.Body)
}
//实现的功能就是curl http://127.0.0.1:8080  -d  '{"Id": "www.361way.com", "Balance": 8}'的功能
