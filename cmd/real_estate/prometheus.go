package main

import (
	"log"
	"strconv"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/push"
)

type PrometheusBackend struct {
    prometheusUrl string
}
func (pb *PrometheusBackend) push(ads []Ad) {
    var (
        price = prometheus.NewGaugeVec(prometheus.GaugeOpts{
                Name: "real_estate_price",
                Help: "The price of the house",
            },
            []string{
                "region",
                "city",
                "link",
                "bathrooms",
                "bedrooms",
                "m2",
            },
        )
        bathrooms = prometheus.NewGaugeVec(prometheus.GaugeOpts{
                Name: "real_estate_bathrooms",
                Help: "The bathrooms of the house",
            },
            []string{
                "region",
                "city",
                "link",
            },
        )
        bedrooms = prometheus.NewGaugeVec(prometheus.GaugeOpts{
                Name: "real_estate_bedrooms",
                Help: "The bedrooms of the house",
            },
            []string{
                "region",
                "city",
                "link",
            },
        )
        m2 = prometheus.NewGaugeVec(prometheus.GaugeOpts{
                Name: "real_estate_m2",
                Help: "The size of the house",
            },
            []string{
                "region",
                "city",
                "link",
            },
        )
    )
    // We use a registry here to benefit from the consistency checks that
    // happen during registration.
    registry := prometheus.NewRegistry()
    registry.MustRegister(price, bathrooms, bedrooms, m2)
    log.Println("Sending to prometheus...")
    for _, ad := range ads {
        price.WithLabelValues(ad.Region, ad.City, ad.Link, strconv.Itoa(ad.Bathrooms), strconv.Itoa(ad.Bedrooms), strconv.Itoa(ad.M2)).Add(ad.Price)
        bathrooms.WithLabelValues(ad.Region, ad.City, ad.Link).Add(float64(ad.Bathrooms))
        bedrooms.WithLabelValues(ad.Region, ad.City, ad.Link).Add(float64(ad.Bedrooms))
        m2.WithLabelValues(ad.Region, ad.City, ad.Link).Add(float64(ad.M2))
        log.Println("sending to prometheus", ad)
    }
    err := push.New(pb.prometheusUrl, "house_market").Gatherer(registry).Push()
    if err != nil {
        log.Fatalln(err)
    }
    log.Println("Done")
}
