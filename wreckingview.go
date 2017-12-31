package wreckingview

import (
    "html/template"
    "net/http"
    
    "wreckingtwo"
    "dataexamples"
)

func init() {
    http.HandleFunc("/dataexamples", dataexamples.ListHandler)
    http.HandleFunc("/dataexamples/save", dataexamples.SaveHandler)
    http.HandleFunc("/dataexamples/upload", dataexamples.UploadHandler)
    http.HandleFunc("/wreckingtwo", wreckingtwo.WreckingtwoHandler)
    http.HandleFunc("/", homeHandler)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    Render(w, "home.html", nil)
}

func Render(w http.ResponseWriter, templateName string, content interface{}) {
    t, _ := template.ParseFiles("layout.html", templateName)
    t.ExecuteTemplate(w, "layout", content)
}