package main 
    import (
        "net/http"
        "fmt"
        "log"
        "encoding/json"
    )
    
     type User struct{
            Id      string
            Balance uint64
     }

func main() {
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                var u User
                if r.Body == nil {
                        http.Error(w, "Please send a request body", 400)
                        return
                }
                err := json.NewDecoder(r.Body).Decode(&u)
                if err != nil {
                        http.Error(w, err.Error(), 400)
                        return
                }
                fmt.Println(u.Id)
        })
        log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
可以使用curl命令进行测试，如下：
curl http://127.0.0.1:8080  -d  '{"Id": "www.361way.com", "Balance": 8}'
如果正常处理得到www.361way.com，就表示功能正常。
*/
