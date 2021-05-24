package main

import "os"
import "log"
import "strings"

func main() {
    queries := strings.Split(os.Getenv("QUERIES"), ",")
    log.Println("getting queries: ", queries)
    getAds(queries)
}
