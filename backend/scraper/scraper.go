package scraper

import (
	"log"
	"vnexpress/models"

	"github.com/gocolly/colly"
	"gorm.io/gorm"
)

func ScrapeAndStore(db *gorm.DB) {
	c := colly.NewCollector(
		colly.AllowedDomains("vnexpress.net"),
	)

	c.OnHTML("article", func(e *colly.HTMLElement) {
		title := e.ChildText("h3.title-news a")
		imageURL := e.ChildAttr("div.thumb-art img", "src")
		content := e.ChildText("p.description")

		//fmt.Println("Title:", title)
		//fmt.Println("Image URL:", imageURL)
		//fmt.Println("Content:", content)

		post := models.Post{Title: title, Image: imageURL, Content: content}
		if err := db.Create(&post).Error; err != nil {
			log.Println("Error inserting data into database:", err)
		}
	})

	c.Visit("https://vnexpress.net/")

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
}
