package main

import "os"
import "strconv"
import "log"

func main() {
    pages, err := strconv.Atoi(os.Getenv("PAGES"))
    if err != nil{
        log.Fatalln(err)
    }
    getAds(pages)
}
