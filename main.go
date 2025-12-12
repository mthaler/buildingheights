package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"

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

	c.OnHTML("table", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			height := el.ChildText("td:nth-child(4)")
			if height != "" {
				height = height[:utf8.RuneCountInString(height)-2]
				height = strings.Replace(height, ",", ".", -1)
			}
			coordinates := el.ChildText("td:nth-child(5)")
			latitude := ""
			longitude := ""
			if coordinates != "" {
				coordinatesArray := strings.Split(coordinates, ",")
				if len(coordinatesArray) >= 2 {
					latitude = coordinatesArray[0]
					latitude = string([]rune(latitude)[:utf8.RuneCountInString(latitude)-2])
					longitude = coordinatesArray[1]
					longitude = string([]rune(longitude)[:utf8.RuneCountInString(longitude)-2])
				}
			}
			b := Building{
				Name:               el.ChildText("td:nth-child(1)"),
				Type:               el.ChildText("td:nth-child(2)"),
				YearOfConstruction: el.ChildText("td:nth-child(3)"),
				Height:             height,
				Latitude:           latitude,
				Longitude:          longitude,
				Remark:             el.ChildText("td:nth-child(6)"),
			}

			w.Write(b.toSlice())

		})
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
