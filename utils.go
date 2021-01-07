// TODO:
// - max 10 requests at the same time
// - update prometheus without processing async
// - Get more ad details - probably

package main

import "log"
import "os"
import "strconv"
import "strings"
import "net/http"
import "github.com/PuerkitoBio/goquery"
import "github.com/prometheus/client_golang/prometheus"
import "github.com/prometheus/client_golang/prometheus/push"


type Ad struct {
    Link string
    City string
    Region string
    Price float64
    Bathrooms int
    Bedrooms int
    M2 int
}

func updateDb(ads []Ad) {
    var (
        price = prometheus.NewGaugeVec(prometheus.GaugeOpts{
                Name: "price",
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
                Name: "bathrooms",
                Help: "The bathrooms of the house",
            },
            []string{
                "region",
                "city",
                "link",
            },
        )
        bedrooms = prometheus.NewGaugeVec(prometheus.GaugeOpts{
                Name: "bedrooms",
                Help: "The bedrooms of the house",
            },
            []string{
                "region",

                "city",
                "link",
            },
        )
        m2 = prometheus.NewGaugeVec(prometheus.GaugeOpts{
                Name: "m2",
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
    log.Println("Sending to prometheus")
    for _, ad := range ads {
        price.WithLabelValues(ad.Region, ad.City, ad.Link, strconv.Itoa(ad.Bathrooms), strconv.Itoa(ad.Bedrooms), strconv.Itoa(ad.M2)).Add(ad.Price)
        bathrooms.WithLabelValues(ad.Region, ad.City, ad.Link).Add(float64(ad.Bathrooms))
        bedrooms.WithLabelValues(ad.Region, ad.City, ad.Link).Add(float64(ad.Bedrooms))
        m2.WithLabelValues(ad.Region, ad.City, ad.Link).Add(float64(ad.M2))
        log.Println("sending to prometheus %s", ad)
        err := push.New(os.Getenv("PROMETHEUS_FQDN"), "house_market").Gatherer(registry).Push()
        if err != nil {
            log.Fatalln(err)
        }
    }
    log.Println("Done")
}

func getGoldenHomePage(page int) {
    log.Printf("Getting page %s...", page)
    query := "https://goldenhome.gr/property/index?PropertySearch%5BPropertyID%5D=&PropertySearch%5BTrnTypeID%5D=2&PropertySearch%5Bvideo_url%5D=&PropertySearch%5BPropCategID%5D=11704&category=&PropertySearch%5BPropSubCategID%5D=&PropertySearch%5BareaLevel1%5D=&PropertySearch%5BRAreaID%5D=&PropertySearch%5BFloorNo%5D=&PropertySearch%5BFloorNo_to%5D=&PropertySearch%5BBuiltYear%5D=1981&PropertySearch%5BBuiltYear_to%5D=&PropertySearch%5BTotalRooms%5D=&PropertySearch%5BTotalRooms_to%5D=&PropertySearch%5BTotalParkings%5D=&PropertySearch%5BTotalParkings_to%5D=&PropertySearch%5BAskedValue%5D=&PropertySearch%5BAskedValue_to%5D=&PropertySearch%5BTotalSm%5D=100&PropertySearch%5BTotalSm_to%5D=&PropertySearch%5Bapothiki%5D=&PropertySearch%5Btzaki%5D=&page=" + strconv.Itoa(page)
    client := &http.Client{}
    req, err := http.NewRequest("GET", query, nil)
    if err != nil {
        log.Println(err)
        return
    }
    req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:77.0) Gecko/20100101 Firefox/77.0")
    resp, err := client.Do(req)
    if err != nil {
        log.Fatalln(err)
    }
    var response []Ad
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
    updateDb(response)
    return
}

func getAds() {
    for i := 1; i < 3000; i++ {
        getGoldenHomePage(i)
	}
    return
}
