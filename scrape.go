package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Blog struct {
	title, content string
}

func main() {
	blogs := ScrapeJDSBlog()
	for _, blog := range blogs {
		fmt.Printf("title:%s, content:%s", blog.title, blog.content)
	}
}

func ScrapeJDSBlog() []Blog {
	// Request the HTML page.
	res, err := http.Get("https://jko.hateblo.jp/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	blogSelection := doc.Find(".entry")
	blogs := []Blog{}

	blogSelection.Each(func(i int, s *goquery.Selection) {
		title := s.Find(".entry-title a").Text()
		var content string
		s.Find(".entry-content > *").Each(func(i int, s *goquery.Selection) {
			if s.Text() != "" {
				content = content + s.Text()
			}
		})
		blogs = append(blogs, Blog{title: title, content: content})
	})

	return blogs
}
