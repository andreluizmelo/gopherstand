package main

import(
    "fmt"
    "net/http"
)

func main() {
    fmt.Println("initializing server...")
    http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request){
        fmt.Println(req)
        w.Write([]byte("Hello Go"))
    })

    http.ListenAndServe(":8000", nil)
}