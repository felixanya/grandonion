package main

import (
    "context"
    "github.com/aaronize/grandonion/examples/basic/context_demo/google"
    "github.com/aaronize/grandonion/examples/basic/context_demo/userip"
    "html/template"
    "log"
    "net/http"
    "time"
)

func main() {
    http.HandleFunc("/search", handleSearch)
    log.Fatal(http.ListenAndServe(":8095", nil))
}
//  GET https://www.googleapis.com/customsearch/v1?key=AIzaSyDUpBLbU7VqQyCJCFnQsl1Vy4fqaEuqLD0&cx=017576662512468239146:omuauf_lfve&q=golang
func handleSearch(w http.ResponseWriter, req *http.Request) {
    var (
        ctx     context.Context
        cancel  context.CancelFunc
    )
    timeout, err := time.ParseDuration(req.FormValue("timeout"))
    if err != nil {
        ctx, cancel = context.WithCancel(context.Background())
    } else {
        ctx, cancel = context.WithTimeout(context.Background(), timeout)
    }
    defer cancel()

    query := req.FormValue("q")
    log.Println("query:", query)
    if query == "" {
        http.Error(w, "no query", http.StatusBadRequest)
    }

    userIp, err := userip.FromRequest(req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    ctx = userip.NewContext(ctx, userIp)

    start := time.Now()
    results, err := google.Search(ctx, query)
    elapsed := time.Since(start)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := resultsTemplate.Execute(w, struct {
        Results     google.Results
        Timeout, Elapsed time.Duration
    }{
        Results: results,
        Timeout: timeout,
        Elapsed: elapsed,
    }); err != nil {
        log.Println(err)
        return
    }
}

var resultsTemplate = template.Must(template.New("results").Parse(`
<html>
<head/>
<body>
  <ol>
  {{range .Results}}
    <li>{{.Title}} - <a href="{{.URL}}">{{.URL}}</a></li>
  {{end}}
  </ol>
  <p>{{len .Results}} results in {{.Elapsed}}; timeout {{.Timeout}}</p>
</body>
</html>
`))
