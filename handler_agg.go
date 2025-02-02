package main

import (
	"fmt"
	"time"
)

func handler_agg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Usage: %v <time_between_reqs>", cmd.name)
	}

	tm, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("Failed to parse time: %v", err)
	}

	ticker := time.NewTicker(tm)
	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			return err
		}
	}

}
