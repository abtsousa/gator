package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("Could not create HTTP request: %v", err)
	}

	req.Header.Set("User-Agent", "gator")

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Could not get HTTP response: %v", err)
	}

	dat, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, fmt.Errorf("Couldn't parse HTTP response: %v", err)
	}

	var rss RSSFeed
	err = xml.Unmarshal(dat, &rss)
	if err != nil {
		return nil, fmt.Errorf("Couldn't unmarshal XML: %v", err)
	}

	rss.Channel.Title = html.UnescapeString(rss.Channel.Title)
	rss.Channel.Description = html.UnescapeString(rss.Channel.Description)

	return &rss, nil

}
