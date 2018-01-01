package wreckingview

import (
    "net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
    Render(w, "views/about.html", nil)
}