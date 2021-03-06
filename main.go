package main

import (
	"encoding/csv"
	"os"
	"time"
	"net/http"
	"fmt"
	"github.com/mmcdole/gofeed"
	"github.com/PuerkitoBio/goquery"
	"gopkg.in/jdkato/prose.v2"
)

func feed(i int) (url string){

	fp := gofeed.NewParser()
	
	feed, err := fp.ParseURL("http://feeds.reuters.com/reuters/UKTopNews")
	
	if err != nil {
		panic(err)
	}

	fmt.Println(feed.Items[i].Title, len(feed.Items))

	url = fmt.Sprintf(feed.Items[i].Link)

	return url

}

func getContent(i int)(content string){
	
	url := feed(i)

	req, err := http.NewRequest("GET",url, nil)
		if err != nil {}

  	client := &http.Client{Timeout: time.Second * 10}
	
	resp, err := client.Do(req)
	if err != nil { fmt.Println("req err")}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
  	if err != nil {}

	s:= fmt.Sprintf(doc.Find(".StandardArticleBody_container").Text())
	
	return s
}

func ent(c chan [][]string){
	for i:= 0; i < 1; i++ {

	s:= getContent(i)

	doc, err := prose.NewDocument(s)
	if err != nil {}

	// Iterate over the doc's named-entities:
    for _, ent := range doc.Entities() {

			data:= [][]string{{ent.Label,ent.Text}}

			c <- data
		}
	}
}

func writetocsv(c chan [][]string){
	file, err := os.Create("tmp.csv")
	if err != nil {}

	writer := csv.NewWriter(file)
    defer writer.Flush()
	 
	for {
		input := <- c
		err := writer.WriteAll(input)
		if err != nil {}
	}
}

func main(){

	var c chan [][]string = make(chan [][]string)

	go ent(c)
	go writetocsv(c)

	var input string
  fmt.Scanln(&input)

}