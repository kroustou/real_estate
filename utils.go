package main

import "log"
import "sync"
import "time"
import "math/rand"
import "strconv"
import "strings"
import "net/http"
import "github.com/PuerkitoBio/goquery"

type Ad struct {
    Link string
    City string
    Region string
    Price float64
    Bathrooms int
    Bedrooms int
    M2 int
}

// TODO: influx/prometheus
func updateDb(ads []Ad) {
    log.Println(ads)
}

func getGoldenHomePage(ch chan []Ad, wg *sync.WaitGroup, page int) {
    time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
    defer wg.Done()
    log.Printf("Getting page %s...", page)
    query := "https://goldenhome.gr/property/index?PropertySearch%5BPropertyID%5D=&PropertySearch%5BTrnTypeID%5D=2&PropertySearch%5Bvideo_url%5D=&PropertySearch%5BPropCategID%5D=11704&category=&PropertySearch%5BPropSubCategID%5D=&PropertySearch%5BareaLevel1%5D=&PropertySearch%5BRAreaID%5D=&PropertySearch%5BFloorNo%5D=&PropertySearch%5BFloorNo_to%5D=&PropertySearch%5BBuiltYear%5D=1981&PropertySearch%5BBuiltYear_to%5D=&PropertySearch%5BTotalRooms%5D=&PropertySearch%5BTotalRooms_to%5D=&PropertySearch%5BTotalParkings%5D=&PropertySearch%5BTotalParkings_to%5D=&PropertySearch%5BAskedValue%5D=&PropertySearch%5BAskedValue_to%5D=&PropertySearch%5BTotalSm%5D=100&PropertySearch%5BTotalSm_to%5D=&PropertySearch%5Bapothiki%5D=&PropertySearch%5Btzaki%5D=&page=" + strconv.Itoa(page)
    client := &http.Client{}
    req, err := http.NewRequest("GET", query, nil)
    if err != nil {
        log.Fatalln(err)
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
        price, _ := strconv.ParseFloat(strings.Split(item.Find(".price").Text(), " €")[0], 64)
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
    ch <- response
}

func getAds() []Ad {
    var wg sync.WaitGroup
    var results []Ad
    ch := make(chan []Ad)
    for i := 1; i < 3000; i++ {
        wg.Add(1)
        go getGoldenHomePage(ch, &wg, i)
	}

    go func() {
        for v := range ch {
            for _, ad := range v {
                results = append(results, ad)
            }
        }
    }()

    wg.Wait()
    close(ch)
    return results
}
