package wreckingview

import (
    "net/http"
)

func init() {
    http.HandleFunc("/data", DataListHandler)
    http.HandleFunc("/data/save", DataSaveHandler)
    http.HandleFunc("/data/upload", UploadHandler)
    http.HandleFunc("/wreckingcreweditor", WreckingcreweditorHandler)
    http.HandleFunc("/about", AboutHandler)
    http.HandleFunc("/track", TrackListHandler)
    http.HandleFunc("/track/save", TrackSaveHandler)
    http.HandleFunc("/track/addonebaby", TrackAddHandler)
    http.HandleFunc("/", HomeHandler)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    Render(w, "views/home.html", nil)
}