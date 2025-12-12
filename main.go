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

	c.OnHTML("table", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			/*height := el.ChildText("td:nth-child(3)")
			if (height != "") {
				height = height[:len(height)-2]
			}
			if err != nil {
				fmt.Printf("Could not convert %s to int")
			}
			s := Summit{ Name: el.ChildText("td:nth-child(2)"),
				 Category: el.ChildText("td:nth-child(4)"),
				 Height: height,
				 Country: el.ChildText("td:nth-child(5)"),
				 Region: el.ChildText("td:nth-child(6)"),
				 Group: el.ChildText("td:nth-child(7)"),
				 Information: el.ChildText("td:nth-child(8)"),
			}
			w.Write(s.toSlice())*/

			b := Building{ Name: el.ChildText("td:nth-child(1)")}

			fmt.Printf("%+v\n", b)
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