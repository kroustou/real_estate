package main

import "os"
import "log"
import "strings"

func main() {
    ag := NewAdGetter()
    queries := strings.Split(os.Getenv("QUERIES"), ",")
    log.Println("getting queries: ", queries)
    ag.getAds(queries)
    ag.updateDb()
}
