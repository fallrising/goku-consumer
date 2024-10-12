package processor

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/fallrising/goku-consumer/internal/api"
	"github.com/fallrising/goku-consumer/pkg/models"
)

func ProcessURL(url string) (*models.URLInfo, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	title := doc.Find("title").First().Text()
	description := doc.Find("meta[name=description]").AttrOr("content", "")

	var tags []string
	doc.Find("meta[name=keywords]").Each(func(i int, s *goquery.Selection) {
		content, _ := s.Attr("content")
		tags = append(tags, strings.Split(content, ",")...)
	})

	return &models.URLInfo{
		URL:         url,
		Title:       title,
		Description: description,
		Tags:        tags,
	}, nil
}

func ProcessMessage(msg []byte, apiClient *api.Client) error {
	content := string(msg)
	urls := extractURLs(content)

	var batch []models.URLInfo
	for _, url := range urls {
		info, err := ProcessURL(url)
		if err != nil {
			log.Printf("Error processing URL %s: %v", url, err)
			continue
		}
		batch = append(batch, *info)

		if len(batch) >= 10 {
			if err := apiClient.UploadBatch(batch); err != nil {
				log.Printf("Error uploading batch: %v", err)
			}
			batch = batch[:0]
		}
	}

	if len(batch) > 0 {
		if err := apiClient.UploadBatch(batch); err != nil {
			log.Printf("Error uploading final batch: %v", err)
		}
	}

	return nil
}

func extractURLs(content string) []string {
	// This is a simple implementation. You might want to use a more robust method,
	// such as a regular expression, to extract URLs.
	words := strings.Fields(content)
	var urls []string
	for _, word := range words {
		if strings.HasPrefix(word, "http://") || strings.HasPrefix(word, "https://") {
			urls = append(urls, word)
		}
	}
	return urls
}
