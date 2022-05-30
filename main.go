package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

func main() {
	fName := "data.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file, err: %q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	_ = writer.Write([]string{"No", "Product Name", "Description", "Image Link", "Price", "Rating", "Name Of Store"})
	defer writer.Flush()

	c := colly.NewCollector(
		colly.AllowedDomains("www.tokopedia.com", "tokopedia.com"),
		colly.UserAgent("xy"),
	)

	detailCollector := c.Clone()

	total := 1
	c.OnHTML(`.css-bk6tzz`, func(e *colly.HTMLElement) {

		detailUrl := e.ChildAttr("a", "href")
		detailUrl = e.Request.AbsoluteURL(detailUrl)
		detailCollector.Visit(detailUrl)
	})

	detailCollector.OnHTML("#main-pdp-container", func(el *colly.HTMLElement) {
		fmt.Println(el.ChildText("h5.css-zeq6c8 > span"))
		writer.Write([]string{
			strconv.Itoa(total),
			el.ChildText(`[data-testid="lblPDPDetailProductName"]`),
			el.ChildText(`[data-testid="lblPDPDescriptionProduk"]`),
			el.ChildAttr(`[data-testid="PDPMainImage"]`, "src"),
			el.ChildText(`[data-testid="lblPDPDetailProductPrice"]`),
			el.ChildText(`[data-testid="lblPDPDetailProductRatingNumber"]`),
			el.ChildText(`[data-testid="llbPDPFooterShopName"]`),
		})

		total++
	})

	fmt.Println("begin scrapping")
	fmt.Println(c.Visit("https://tokopedia.com/p/handphone-tablet/handphone"))
	fmt.Println(c.Visit("https://tokopedia.com/p/handphone-tablet/handphone?page=2"))
	fmt.Println(c.Visit("https://tokopedia.com/p/handphone-tablet/handphone?page=3"))
	fmt.Println(c.Visit("https://tokopedia.com/p/handphone-tablet/handphone?page=4"))
	fmt.Println(c.Visit("https://tokopedia.com/p/handphone-tablet/handphone?page=5"))
	fmt.Println(c.Visit("https://tokopedia.com/p/handphone-tablet/handphone?page=6"))
	fmt.Println(c.Visit("https://tokopedia.com/p/handphone-tablet/handphone?page=7"))
	fmt.Println(c.Visit("https://tokopedia.com/p/handphone-tablet/handphone?page=8"))
	fmt.Println(c.Visit("https://tokopedia.com/p/handphone-tablet/handphone?page=9"))
	fmt.Println(c.Visit("https://tokopedia.com/p/handphone-tablet/handphone?page=10"))
}
