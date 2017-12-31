package wreckingtwo

import (
    "net/http"
    
    "common"
)

func WreckingtwoHandler(w http.ResponseWriter, r *http.Request) {
    common.Render(w, "wreckingtwo/wreckingtwo.html", nil)
}