package wreckingview

import (
    "net/http"
)

func WreckingtwoHandler(w http.ResponseWriter, r *http.Request) {
    Render(w, "views/wreckingtwo.html", nil)
}