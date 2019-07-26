package google

import (
    "context"
    "encoding/json"
    "github.com/aaronize/grandonion/examples/basic/context_demo/userip"
    "io/ioutil"
    "log"
    "net/http"
)

type Results []Result

type Result struct {
    Title, URL string
}

func Search(ctx context.Context, query string) (Results, error) {
    req, err := http.NewRequest("GET", "https://www.googleapis.com/customsearch/v1?key=AIzaSyDUpBLbU7VqQyCJCFnQsl1Vy4fqaEuqLD0&cx=017576662512468239146:omuauf_lfve", nil)
    if err != nil {
        return nil, err
    }
    q := req.URL.Query()
    q.Set("q", query)

    if userIp, ok := userip.FromContext(ctx); ok {
        q.Set("userip", userIp.String())
    }
    req.URL.RawQuery = q.Encode()

    var results Results
    err = httpDo(ctx, req, func(resp *http.Response, err error) error {
        if err != nil {
            return err
        }
        defer resp.Body.Close()
        by, _ := ioutil.ReadAll(resp.Body)
        log.Println(string(by))

        var data struct{
            responseData struct{
                Results []struct{
                    TitleNotFormatting  string
                    URL                 string
                }
            }
        }

        if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
            return err
        }
        for _, res := range data.responseData.Results {
            results = append(results, Result{Title: res.TitleNotFormatting, URL: res.URL})
        }
        return nil
    })

    return results, err
}

func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
    c := make(chan error, 1)
    req = req.WithContext(ctx)
    go func() { c <- f(http.DefaultClient.Do(req)) }()
    select {
    case <- ctx.Done():
        <-c
        return ctx.Err()
    case err := <-c:
        return err
    }
}
