package main

import (
	"context"
	"fmt"
)

func handler_feeds(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Usage: %v", cmd.name)
	}

	fds, err := s.db.GetFeeds(context.Background())

	if err != nil {
		return fmt.Errorf("Couldn't retrieve feeds: %v", err)
	}

  for _, fd := range fds {
    fmt.Printf("* %s (%s) [added by %s]\n", fd.Name, fd.Url, fd.User)
  }

	return nil
}
