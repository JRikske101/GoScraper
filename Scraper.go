package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/gocolly/colly"
)

type item struct {
	Name   string `json: "name"`
	Price  string `json: "price"`
	ImgUrl string `json: "imgUrl"`
}

func main() {
	// fileName := "GoScraper.csv"
	// file, err = os.Create(fileName)

	// if err != nil {
	// 	log.Fatalf("Cannot create file %q: %s: \n", fileName, err)
	// 	return
	// }
	// defer file.close()
	// writer := csv.NewWriter(fileName)
	// defer writer.Flush()

	c := colly.NewCollector(
	//colly.AllowedDomains("bol.com"),
	)
	c.DetectCharset = true
	c.SetRequestTimeout(120 * time.Second)

	c.OnHTML("div.columns .column.main #amasty-shopby-product-list .products .products .item .product-item-info .product", func(h *colly.HTMLElement) {

		var items []item

		item := item{
			Name:   h.ChildText(".tbl"),
			Price:  h.ChildText(".price-wrapper"),
			ImgUrl: h.ChildAttr("img", "src"),
		}

		items = append(items, item)
		items = delete_empty(items)

		fmt.Println(items)

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response: ", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error: ", e)
	})

	c.Visit("https://www.petsplace.nl/hond/hondensnack")
}

func delete_empty(s []item) []item {
	var r []item
	for _, item := range s {
		if !reflect.ValueOf(item).IsZero() {
			r = append(r, item)
		}
	}
	return r
}
