package main

import (
	"context"
	"fmt"

	"github.com/abtsousa/gator/internal/database"
	"github.com/google/uuid"
)

func handler_follow(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Usage: %v <urL>", cmd.name)
	}

	url := cmd.args[0]

	fd, err := s.db.GetFeed(context.Background(), url)

	if err != nil {
		return fmt.Errorf("Couldn't retrieve feed: %v", err)
	}

	user, _:= s.db.GetUser(context.Background(), s.cfg.CurrentUserName)

	_, err = s.db.CreateFeedFollow(context.Background(),
	database.CreateFeedFollowParams {
		ID: uuid.New(),
		CreatedAt: fd.CreatedAt,
		UpdatedAt: fd.UpdatedAt,
		UserID: user.ID,
		FeedID: fd.ID,
		})

	if err != nil {
		return fmt.Errorf("Failed to follow feed.", err)
	}

	return nil
}
