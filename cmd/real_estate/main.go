package main

import "os"
import "log"
import "strings"

func main() {
    ag := AdGetter{prometheusUrl: os.Getenv("PROMETHEUS_FQDN")}
    queries := strings.Split(os.Getenv("QUERIES"), ",")
    log.Println("getting queries: ", queries)
    ag.getAds(queries)
}
