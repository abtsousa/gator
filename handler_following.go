package main

import (
	"context"
	"fmt"

	"github.com/abtsousa/gator/internal/database"
)

func handler_following(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Usage: %v", cmd.name)
	}

	fds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)

	if err != nil {
		return fmt.Errorf("Couldn't retrieve feeds: %v", err)
	}

  for _, fd := range fds {
    fmt.Printf("* %s\n", fd.Name)
  }

	return nil
}
