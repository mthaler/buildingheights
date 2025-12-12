package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

func main() {
	fName := "buildings.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	c := colly.NewCollector()

 	// Define the URL you want to scrape
 	url := "https://de.wikipedia.org/wiki/Liste_der_h%C3%B6chsten_Bauwerke_in_M%C3%BCnchen"

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Scraping:", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
       	fmt.Println("Status:", r.StatusCode)
   	})

	c.OnError(func(r *colly.Response, err error) {
    	fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
   	})

	// Visit the URL and start scraping
	err = c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
}