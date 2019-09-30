package configure

import (
    "fmt"
    "log"
    "testing"
)

func TestParseConfig(t *testing.T) {
    conf, err := ParseConfig("./config.toml")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("---> %+v\n", conf)
}
