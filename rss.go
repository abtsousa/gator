package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/abtsousa/gator/internal/database"
	"github.com/google/uuid"
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

func scrapeFeeds(s *state) error {
	fd, err := s.db.GetNextFeedToFetch(context.Background())

	if err != nil {
		return fmt.Errorf("Couldn't retrieve feeds: %v", err)
	}

	rss, err := fetchFeed(context.Background(), fd.Url)
	if err != nil {
		return err
	}

    err = s.db.MarkFeedFetched(context.Background(), fd.ID)
    if err != nil {
            return fmt.Errorf("Couldn't mark feed as fetched: %v", err)
    }

    for _, item := range rss.Channel.Item {

		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}
    	
    	_, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
    		ID:          uuid.New(),
    		CreatedAt:   time.Now().UTC(),
    		UpdatedAt:   time.Now().UTC(),
    		Title:       sql.NullString{String: item.Title, Valid: true},
    		Url:         item.Link,
    		Description: sql.NullString{String: item.Description, Valid: true},
    		PublishedAt: publishedAt,
    		FeedID:      fd.ID,
    	})
    	if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
    	}
  	}

	return nil
}
