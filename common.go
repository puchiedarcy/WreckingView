package wreckingview

import (
    "html/template"
    "net/http"
)

func Render(w http.ResponseWriter, templateName string, content interface{}) {
    t, _ := template.ParseFiles("views/layout.html", templateName)
    t.ExecuteTemplate(w, "layout", content)
}