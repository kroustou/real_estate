package main

import "log"
import "strings"
import "fmt"
import "net/http"
import "github.com/PuerkitoBio/goquery"


func getAds() {
    query := "https://goldenhome.gr/property/index?PropertySearch%5BPropertyID%5D=&PropertySearch%5BTrnTypeID%5D=2&PropertySearch%5Bvideo_url%5D=&PropertySearch%5BPropCategID%5D=11704&category=&PropertySearch%5BPropSubCategID%5D=&PropertySearch%5BareaLevel1%5D=&PropertySearch%5BRAreaID%5D=&PropertySearch%5BFloorNo%5D=&PropertySearch%5BFloorNo_to%5D=&PropertySearch%5BBuiltYear%5D=1981&PropertySearch%5BBuiltYear_to%5D=&PropertySearch%5BTotalRooms%5D=&PropertySearch%5BTotalRooms_to%5D=&PropertySearch%5BTotalParkings%5D=&PropertySearch%5BTotalParkings_to%5D=&PropertySearch%5BAskedValue%5D=&PropertySearch%5BAskedValue_to%5D=&PropertySearch%5BTotalSm%5D=100&PropertySearch%5BTotalSm_to%5D=&PropertySearch%5Bapothiki%5D=&PropertySearch%5Btzaki%5D="
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

    defer resp.Body.Close()
    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }

    doc.Find(".pgl-property").Each(func(index int, item *goquery.Selection) {
        address, _ := item.Find("address").Html()
        brokenDownAddress := strings.Split(address, "<br/>")
        city := brokenDownAddress[0]
        region := brokenDownAddress[1]
        fmt.Println(city)
        fmt.Println(region)
        fmt.Println(item.Find(".price").Text())
        fmt.Println(strings.TrimSpace(item.Find(".amenities .pull-left").Text())[8:])
        fmt.Println(strings.TrimSpace(item.Find(".amenities .pull-right").Text()))
    })
}
