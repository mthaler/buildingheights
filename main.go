package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

func main() {
	fName := "summits.csv"
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

}