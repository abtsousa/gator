package main

import (
	"context"
	"fmt"
)

func handler_following(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Usage: %v", cmd.name)
	}

	user, _:= s.db.GetUser(context.Background(), s.cfg.CurrentUserName)


	fds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)

	if err != nil {
		return fmt.Errorf("Couldn't retrieve feeds: %v", err)
	}

  for _, fd := range fds {
    fmt.Printf("* %s\n", fd.Name)
  }

	return nil
}
