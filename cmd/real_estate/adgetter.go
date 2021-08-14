package main

import (
    "os"
    "log"
)

type Ad struct {
    Link string
    City string
    Region string
    Price float64
    Bathrooms int
    Bedrooms int
    M2 int
}

type StorageBackend interface {
    push([]Ad)
}

type AdBackend interface {
    getQuery(string) []Ad
}

type AdGetter struct {
    ads []Ad
    backend AdBackend
    storageBackend StorageBackend
}


func (ag *AdGetter) updateDb() {
    ag.storageBackend.push(ag.ads)
}


func (ag *AdGetter) getAds(queries []string) {
    log.Println("Getting data")
    // TODO Concurrency
    for _, query := range queries {
        log.Println("query", query)
        ag.ads = append(ag.ads, ag.backend.getQuery(query)...)
        log.Println(ag.ads)
    }
}

func NewAdGetter() AdGetter {
    return AdGetter{
        storageBackend: &PrometheusBackend{prometheusUrl: os.Getenv("PROMETHEUS_FQDN")},
        backend: &GoldenHomeBackend{},
    }
}
