package main 
    import (
        "net/http"
        "log"
        "encoding/json"
    )

 type User struct{
          Id      string
          Balance uint64
   }
func main() {
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                u := User{Id: "www.361way.com", Balance: 8}
                json.NewEncoder(w).Encode(u)
        })
        log.Fatal(http.ListenAndServe(":8080", nil))
}
/*
通过curl http://127.0.0.1:8080 其会返回json数据给客户端
*/
