package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
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
			heightText := el.ChildText("td:nth-child(4)")
			if heightText != "" {
				heightText = heightText[:utf8.RuneCountInString(heightText)-2]
				heightText = strings.Replace(heightText, ",", ".", -1)
				height, err := strconv.ParseFloat(heightText, 64)
				if err != nil {
					fmt.Printf("Error:%v\n", err)
				}
				height = math.Round(height)
				heightInt := int(height)
				heightText := fmt.Sprintf("%d", heightInt)
				fmt.Printf("%s\n", heightText)
			}
			b := Building{
				Name:               el.ChildText("td:nth-child(1)"),
				Type:               el.ChildText("td:nth-child(2)"),
				YearOfConstruction: el.ChildText("td:nth-child(3)"),
				Height:             heightText,
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
