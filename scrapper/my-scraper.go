package main

import (
	"fmt"
	"log"
	"strings"
	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()
	var name string
	fmt.Println("enter the registered name of the comapny")
	fmt.Scan(&name)
	url := "https://finance.yahoo.com/quote/"+name
	c.OnHTML(".livePrice",func (e *colly.HTMLElement)  {
		pricing := e.ChildText("span")
		price := strings.TrimSpace(pricing)
		fmt.Println((price))
	})
   
	err := c.Visit(url)
	if err != nil {
	 log.Fatal(err)
	}
   }