package wreckingtwo

import (
    "fmt"
    "html/template"
    "net/http"
)

func WreckingtwoHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprint(w, "<!DOCTYPE html><html>")
    fmt.Fprint(w, "<head><meta charset='UTF-8'></head>")
    fmt.Fprint(w, "<body>")
    
    t2, _ := template.ParseFiles("wreckingtwo/wreckingtwo.html")
    t2.Execute(w, nil)
    
    fmt.Fprint(w, "</body></html>")
}