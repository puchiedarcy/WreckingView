package wreckingview

import (
    "net/http"
    
    "wreckingtwo"
    "dataexamples"
    "common"
)

func init() {
    http.HandleFunc("/dataexamples", dataexamples.ListHandler)
    http.HandleFunc("/dataexamples/save", dataexamples.SaveHandler)
    http.HandleFunc("/dataexamples/upload", dataexamples.UploadHandler)
    http.HandleFunc("/wreckingtwo", wreckingtwo.WreckingtwoHandler)
    http.HandleFunc("/", homeHandler)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    common.Render(w, "home.html", nil)
}