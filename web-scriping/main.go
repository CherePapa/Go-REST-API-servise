package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Product struct {
	Url, Image, Name, Price string
}

func main() {

	c := colly.NewCollector(
		colly.AllowedDomains("www.scrapingcourse.com"),
	)

	var products []Product

	c.OnHTML("li.product", func(e *colly.HTMLElement) {

		product := Product{}

		product.Url = e.ChildAttr("a", "href")
		product.Image = e.ChildAttr("img", "scr")
		product.Name = e.ChildText(".product-name")
		product.Price = e.ChildText(".price")

		products = append(products, product)
	})

	c.OnScraped(func(r *colly.Response) {
		file, err := os.Create("product.csv")

		if err != nil {
			log.Fatalln("Failed to create output CSV File", err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)

		headers := []string{
			"Url",
			"Image",
			"Name",
			"Price",
		}
		writer.Write(headers)

		for _, product := range products {
			record := []string{
				product.Url,
				product.Image,
				product.Name,
				product.Price,
			}

			writer.Write(record)
		}

		defer writer.Flush()
	})

	c.Visit("https://www.scrapingcourse.com/ecommerce")
}
