package main

import (
    "log"
    "fmt"
    "os"
    "strconv"
    "strings"
    "net/http"
    "github.com/PuerkitoBio/goquery"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/push"
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

type AdGetter struct {
    ads []Ad
    prometheusUrl string
}

func (ag *AdGetter) updateDb() {
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
    for _, ad := range ag.ads {
        price.WithLabelValues(ad.Region, ad.City, ad.Link, strconv.Itoa(ad.Bathrooms), strconv.Itoa(ad.Bedrooms), strconv.Itoa(ad.M2)).Add(ad.Price)
        bathrooms.WithLabelValues(ad.Region, ad.City, ad.Link).Add(float64(ad.Bathrooms))
        bedrooms.WithLabelValues(ad.Region, ad.City, ad.Link).Add(float64(ad.Bedrooms))
        m2.WithLabelValues(ad.Region, ad.City, ad.Link).Add(float64(ad.M2))
        log.Println("sending to prometheus", ad)
    }
    err := push.New(os.Getenv("PROMETHEUS_FQDN"), "house_market").Gatherer(registry).Push()
    if err != nil {
        log.Fatalln(err)
    }
    log.Println("Done")
}


func (ag *AdGetter) getGoldenHomePage(page int, query string) []Ad {
    log.Printf("Getting page %d...", page)
    var response []Ad
    client := &http.Client{}
    req, err := http.NewRequest("GET", fmt.Sprintf("%s&page=%d", query, page), nil)
    if err != nil {
        log.Fatalln(err)
    }
    req.Header.Set("User-Agent", "real_estate bot")
    resp, err := client.Do(req)
    if err != nil {
        log.Fatalln(err)
    }
    defer resp.Body.Close()
    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }
    doc.Find(".pgl-property").Each(func(index int, item *goquery.Selection) {
        link, _ := item.Find("a").Attr("href")
        address, _ := item.Find("address").Html()
        brokenDownAddress := strings.Split(address, "<br/>")
        city := brokenDownAddress[0]
        region := ""
        if (len(brokenDownAddress) == 2) {
            region = brokenDownAddress[1]
        }
        price, _ := strconv.ParseFloat(strings.Replace(strings.Split(item.Find(".price").Text(), " €")[0], ".", "", -1), 64)
        m2, _ := strconv.Atoi(strings.Split(strings.TrimSpace(item.Find(".amenities .pull-left").Text())[8:], " ")[0])
        var bedroomsAndBathrooms [2]int
        item.Find(".amenities .pull-right li").Each(func(index int, item *goquery.Selection){
            bedroomsAndBathrooms[index], _ = strconv.Atoi(strings.TrimSpace(item.Text()))
        })
        bedrooms := bedroomsAndBathrooms[0]
        bathrooms := bedroomsAndBathrooms[1]
        response = append(response, Ad{
            Link: "https://goldenhome.gr" + link,
            City: city,
            Region: region,
            Price: price,
            Bathrooms: bathrooms,
            Bedrooms: bedrooms,
            M2: m2,
        })
    })
    log.Println("Got ", len(response))
    return response
}


func (ag *AdGetter) getAds(queries []string) {
    log.Println("Getting data")
    for _, query := range queries {
        log.Println("query", query)
        // a flag on whether we should fetch the next page
        // goldenhome will return the last page as soon as you have exceeded the pages
        // and since the pagination relies on infinite scrolling there is no other way to see
        // whether we have reached the end
        var getNext bool = true
        var lastLink string
        page := 1
        for getNext {
            response := ag.getGoldenHomePage(page, query)
            page += 1
            // We just fetched a page that has been already stored
            if lastLink == response[len(response)-1].Link {
                break
            }
            // else we need to add the new items
            lastLink = response[len(response)-1].Link
            for r := 0; r < len(response); r++ {
                // if item already in list, stop
                ag.ads = append(ag.ads, response...)
            }
        }
        log.Println(ag.ads)
    }
    if len(ag.ads) > 0 {
        ag.updateDb()
    }
}
