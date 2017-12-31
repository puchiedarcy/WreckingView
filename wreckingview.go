package wreckingview

import (
    "fmt"
    "net/http"
    
    "wreckingtwo"
    "dataexamples"
)

func init() {
    http.HandleFunc("/favicon.ico", faviconHandler)
    http.HandleFunc("/dataexamples", dataexamples.ListHandler)
    http.HandleFunc("/dataexamples/save", dataexamples.SaveHandler)
    http.HandleFunc("/dataexamples/upload", dataexamples.UploadHandler)
    http.HandleFunc("/wreckingtwo", wreckingtwo.WreckingtwoHandler)
    http.HandleFunc("/", wreckingviewHandler)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
    return
}

func wreckingviewHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprint(w, "<!DOCTYPE html><html>")
    fmt.Fprint(w, "<head><meta charset='UTF-8'></head>")
    fmt.Fprint(w, "<body>")
    fmt.Fprint(w, "<h3>Welcome to Wrecking View!</h3>")
    
    fmt.Fprint(w, "<a href='/dataexamples'>Data Examples</a><br>")
    fmt.Fprint(w, "<a href='/wreckingtwo'>Wrecking Two</a>")
    
    fmt.Fprint(w, "</body></html>")
}