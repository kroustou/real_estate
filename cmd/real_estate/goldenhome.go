package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type GoldenHomeBackend struct{}

func (ghb *GoldenHomeBackend) getQuery(query string) []Ad {
	// a flag on whether we should fetch the next page
	// goldenhome will return the last page as soon as you have exceeded the pages
	// and since the pagination relies on infinite scrolling there is no other way to see
	// whether we have reached the end
	var getNext bool = true
	var lastLink string
	var ads []Ad
	page := 1
	for getNext {
		response := ghb.getPage(page, query)
		page += 1
		// We just fetched a page that has been already stored
		if lastLink == response[len(response)-1].Link {
			break
		}
		// else we need to add the new items
		lastLink = response[len(response)-1].Link
		for r := 0; r < len(response); r++ {
			// if item already in list, stop
			ads = append(ads, response...)
		}
	}
	return ads
}

func (ghb *GoldenHomeBackend) getPage(page int, query string) []Ad {
	log.Printf("Getting page %d...", page)
	var url = fmt.Sprintf("%s&page=%d", query, page)
	return ghb.getAdsUrl(url)
}

func (ghb *GoldenHomeBackend) getAdsUrl(url string) []Ad {
	var response []Ad
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
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
		if len(brokenDownAddress) == 2 {
			region = brokenDownAddress[1]
		}
		price, _ := strconv.ParseFloat(strings.Replace(strings.Split(item.Find(".price").Text(), " €")[0], ".", "", -1), 64)
		m2, _ := strconv.Atoi(strings.Split(strings.TrimSpace(item.Find(".amenities .pull-left").Text())[8:], " ")[0])
		var bedroomsAndBathrooms [2]int
		item.Find(".amenities .pull-right li").Each(func(index int, item *goquery.Selection) {
			bedroomsAndBathrooms[index], _ = strconv.Atoi(strings.TrimSpace(item.Text()))
		})
		bedrooms := bedroomsAndBathrooms[0]
		bathrooms := bedroomsAndBathrooms[1]
		response = append(response, Ad{
			Link:      "https://goldenhome.gr" + link,
			City:      city,
			Region:    region,
			Price:     price,
			Bathrooms: bathrooms,
			Bedrooms:  bedrooms,
			M2:        m2,
		})
	})
	log.Println("Got ", len(response))
	return response
}
