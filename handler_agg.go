package main

import (
	"context"
	"fmt"
)

func handler_agg(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Usage: %v", cmd.name)
	}

	rss, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Println(rss)

	return nil
}
