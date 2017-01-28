package main

import(
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"
)

func main() {
    fmt.Println("initializing server...")
    // http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request){
    //     fmt.Println(req)
    //     w.Write([]byte("Hello Go"))
    // })
    http.Handle("/", new(MyHandler))

    http.ListenAndServe(":8000", nil)
}

type MyHandler struct {
    http.Handler
}

func (this *MyHandler) ServeHTTP( w http.ResponseWriter, req *http.Request){
    path := "./" + req.URL.Path
    data, error := ioutil.ReadFile(string(path))
    if error == nil {
        var contentType string
        if strings.HasSuffix(path, ".css"){
            contentType = "text/css"
        } else if strings.HasSuffix(path, ".html"){
            contentType = "text/html"
        } else if strings.HasSuffix(path, ".js"){
            contentType = "application/javascript"
        } else if strings.HasSuffix(path, ".png"){
            contentType = "image/png"
        } else {
            contentType = "text/plain"
        }
        w.Header().Add("Content Type", contentType)
        w.Write(data)
    } else {
        w.WriteHeader(404)
        w.Write([]byte("404 - " + http.StatusText(404)))
    }
}