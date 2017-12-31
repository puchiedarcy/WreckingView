package wreckingview

import (
    "fmt"
    "html/template"
    "net/http"
    "io"
    
    "google.golang.org/appengine"
    "google.golang.org/appengine/datastore"
    "google.golang.org/appengine/file"
    "cloud.google.com/go/storage"
)

type Phase struct {
    Name string
    Number string
}

type List struct {
    Headers []string
    Rows    []interface{}
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
    ctx := appengine.NewContext(r)
    
    p := Phase {
        Number: r.FormValue("Number"),
        Name: r.FormValue("Name"),
    }
    
    _, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "Phase", nil), &p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    http.Redirect(w, r, "/dataexamples", http.StatusSeeOther)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
    ctx := appengine.NewContext(r)
    
    client, err := storage.NewClient(ctx)
    bucketName, _ := file.DefaultBucketName(ctx)
    bucket := client.Bucket(bucketName)
    
    f, fh, err := r.FormFile("image")
    if err == http.ErrMissingFile {
        return
    }
    if err != nil {
        return
    }
    
    out := bucket.Object(fh.Filename).NewWriter(ctx)
    
    if _, err := io.Copy(out, f); err != nil {
        fmt.Fprint(w, "p1 ", err)
    }
    if err := out.Close(); err != nil {
        fmt.Fprint(w, "P2 ", err)
    }
    
    http.Redirect(w, r, "/dataexamples", http.StatusSeeOther)
}


func ListHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprint(w, "<!DOCTYPE html><html>")
    fmt.Fprint(w, "<head><meta charset='UTF-8'></head>")
    fmt.Fprint(w, "<body>")
    fmt.Fprint(w, "<a href='/'>Back</a><br>")
    ctx := appengine.NewContext(r)
    
    t2, _ := template.ParseFiles("save.html")
    t2.Execute(w, nil)
    
    q := datastore.NewQuery("Phase").Order("Number")
    count, _ := q.Count(ctx)
    
    list := List {
        Headers: []string{"Number", "Name", "Image"},
        Rows: make([]interface{}, count),
    }
    
    t := q.Run(ctx)
    i := 0
    for {
        var p Phase
        _, err := t.Next(&p)
        if (err == datastore.Done) {
            break
        }
        
        list.Rows[i] = p
        i++
    }
    
    tl, _ := template.ParseFiles("list.html")
    tl.ExecuteTemplate(w, "list", list)
    
    fmt.Fprint(w, "</body></html>")
}