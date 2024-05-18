package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=python"

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

func main() {
	totalPages := getPages()
	fmt.Println(totalPages)

	for i := 0; i < totalPages; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	pageUrl := baseURL + "&recruitPage=" + strconv.Itoa(page)
	fmt.Println("Requesting ", pageUrl)
	res, err := http.Get(pageUrl)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".item_recruit")

	searchCards.Each(func(i int, card *goquery.Selection) {
		extractJob(card)
	})

}

func extractJob(card *goquery.Selection) {
	id, _ := card.Attr("value")
	fmt.Println(id)
	title := card.Find(".job_tit>a").Text()
	fmt.Println(title)
	location := card.Find(".job_condition> span").First().Text()
	fmt.Println(location)
}

func getPages() int {

	// req, err := http.NewRequest("GET", baseURL, nil)
	// checkErr(err)

	// purl, err := url.Parse(baseURL)
	// checkErr(err)

	// client := &http.Client{
	// 	Transport: &http.Transport{Proxy: http.ProxyURL(purl)}}

	// res, err := client.Do(req)
	// checkErr(err)
	// checkCode(res)

	// defer res.Body.Close()
	pages := 0

	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	// fmt.Println(doc)
	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("request failed with status", res.StatusCode)
	}
}
