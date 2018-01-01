package wreckingview

import (
    "net/http"
)

func WreckingcreweditorHandler(w http.ResponseWriter, r *http.Request) {
    Render(w, "views/wreckingcreweditor.html", nil)
}