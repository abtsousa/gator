package main

import (
	"context"
	"fmt"
	"time"

	"github.com/abtsousa/gator/internal/database"
	"github.com/google/uuid"
)

func handler_addfeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("Usage: %v <name> <url>", cmd.name)
	}

	name, url := cmd.args[0], cmd.args[1]

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Couldn't retrieve current user: %v", err)
	}

	fd, err := s.db.AddFeed(context.Background(), database.AddFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("Couldn't add feed: %v", err)
	}


	fmt.Println("Added new feed.")

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
