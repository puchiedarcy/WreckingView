package wreckingview

import (
    "net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    Render(w, "views/login.html", nil)
}