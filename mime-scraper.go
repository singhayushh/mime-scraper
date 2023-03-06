package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly"
)

type FileType struct {
	Name      string
	Mimetype  string
	Icon      string
	Extension string
	Info      string
}

func main() {
	c := colly.NewCollector()
	c.SetRequestTimeout(10 * time.Second)

	fileTypes := make([]FileType, 0)

	c.OnHTML("tbody", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			fileType := FileType{}
			fileType.Name = el.ChildText("td:nth-child(1)")
			fileType.Mimetype = el.ChildText("td:nth-child(2)")
			fileType.Icon = ""
			fileType.Extension = el.ChildText("td:nth-child(3)")
			fileType.Info = el.ChildAttr("td:nth-child(4) a", "href")
			fileTypes = append(fileTypes, fileType)
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Scraping:", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response:", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		js, err := json.MarshalIndent(fileTypes, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Writing data to file")
		if err := os.WriteFile("mimetypes.json", js, 0664); err == nil {
			fmt.Println("Data written to file successfully")
		}
	})

	c.Visit("https://www.freeformatter.com/mime-types-list.html")
}
