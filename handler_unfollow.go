package main

import (
	"context"
	"fmt"

	"github.com/abtsousa/gator/internal/database"
)

func handler_unfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Usage: %v <feed>", cmd.name)
	}

	f := cmd.args[0]

	fd, err := s.db.GetFeed(context.Background(), f)

	if err != nil {
		return fmt.Errorf("Couldn't retrieve feed: %v", err)
	}

	err = s.db.DeleteFeedFollow(context.Background(),
		database.DeleteFeedFollowParams{UserID: user.ID, FeedID: fd.ID})

	if err != nil {
		return fmt.Errorf("Failed to unfollow feed: %v", err)
	}

	fmt.Println("Unfollowed from feed.")
	return nil
}
