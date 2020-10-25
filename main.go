package main

import "fmt"
import "log"
import "io/ioutil"
import "net/http"

func getAds(query string) string {
    resp, err := http.Get("https://xe.gr/" + query)
    if err != nil {
        log.Fatal(err)
    }
    body, err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    if err != nil {
        log.Fatal(err)
    }
    return string(body)
}

func main() {
    fmt.Println(getAds("?"))
}
