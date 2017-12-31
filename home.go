package wreckingview

import (
    "net/http"
)

func init() {
    http.HandleFunc("/data", ListHandler)
    http.HandleFunc("/data/save", SaveHandler)
    http.HandleFunc("/data/upload", UploadHandler)
    http.HandleFunc("/wreckingtwo", WreckingtwoHandler)
    http.HandleFunc("/", HomeHandler)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    Render(w, "views/home.html", nil)
}