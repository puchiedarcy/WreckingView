package wreckingview

import (
    "time"
    "net/http"
    "fmt"
    "strings"
    
    "google.golang.org/appengine"
    "google.golang.org/appengine/datastore"
)

type Track struct {
    User string
    Name string
    Score int
    HighScore int
    DateCreated time.Time
    DateUpdated time.Time
    Public bool
}

type List struct {
    Headers []string
    Rows    []interface{}
}

func TrackListHandler(w http.ResponseWriter, r *http.Request) {
    ctx := appengine.NewContext(r)
    
    q := datastore.NewQuery("Track").Order("DateCreated")
    count, _ := q.Count(ctx)
    
    list := List {
        Headers: []string{"Name", "Score", "HighScore", "DateCreated", "DateUpdated", "Public"},
        Rows: make([]interface{}, count),
    }
    
    qr := q.Run(ctx)
    i := 0
    for {
        var t Track
        k, err := qr.Next(&t)
        if (err == datastore.Done) {
            break
        }
        
        est, _ := time.LoadLocation("EST")
        data := [6]string{
            t.Name,
            fmt.Sprintf("%d", t.Score),
            fmt.Sprintf("%d", t.HighScore),
            t.DateCreated.In(est).Format("Mon Jan 2 2006"),
            t.DateUpdated.In(est).Format("Mon Jan 2 2006"),
            fmt.Sprintf("%t", t.Public),
        }
        
        viewData := struct {
            Key string
            Data [6]string
        }{
            strings.SplitN(fmt.Sprintf("%s", k), ",", 2)[1],
            data,
        }
        
        list.Rows[i] = viewData
        i++
    }
    
    Render(w, "views/track.html", list)
}

func TrackSaveHandler(w http.ResponseWriter, r *http.Request) {
    ctx := appengine.NewContext(r)
    
    t := Track {
        User: "puchie",
        Name: r.FormValue("Name"),
        Score: 0,
        HighScore: 0,
        DateCreated: time.Now(),
        DateUpdated: time.Now(),
        Public: r.FormValue("Public") == "public",
    }
   
    _, err := datastore.Put(ctx, datastore.NewKey(ctx, "Track", "puchie-" + t.Name, 0, nil), &t)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    http.Redirect(w, r, "/track", http.StatusSeeOther)
}

func TrackAddHandler(w http.ResponseWriter, r *http.Request) {
    ctx := appengine.NewContext(r)

    k := datastore.NewKey(ctx, "Track", r.FormValue("Key"), 0, nil)
    var t Track
    datastore.Get(ctx, k, &t)
    
    t.Score++
    
    datastore.Put(ctx, k, &t)
    
    http.Redirect(w, r, "/track", http.StatusSeeOther)
}