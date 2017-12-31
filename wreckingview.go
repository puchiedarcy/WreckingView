package wreckingview

import (
    "net/http"
)

func init() {
    http.HandleFunc("/dataexamples", ListHandler)
    http.HandleFunc("/dataexamples/save", SaveHandler)
    http.HandleFunc("/dataexamples/upload", UploadHandler)
    http.HandleFunc("/wreckingtwo", WreckingtwoHandler)
    http.HandleFunc("/", homeHandler)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    Render(w, "home.html", nil)
}