package url

import (
    "testing"
    "fmt"
)

func TestUrlSplit(t *testing.T) {
    // url := UrlSplit("http://127.0.0.1:50000/eureka?name=Jake&age=20")
    url := UrlSplit("http://127.0.0.1:50000/index?user=Jake&pwd=123456")
    fmt.Println(url.Proto.String(), url.Addr.String(), url.Path.String())
    for _, v := range url.Params {
        fmt.Println(v.Key.String(), v.Value.String())
    }
}
