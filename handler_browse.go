package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/abtsousa/gator/internal/database"
)

func handler_browse(s *state, cmd command, user database.User) error {

	var limit int
	switch l:= len(cmd.args); {
		case l > 1:
			return fmt.Errorf("Usage: %v <number>", cmd.name)
		case l == 1:
			lmt, err := strconv.Atoi(cmd.args[0])
			if err != nil {
				return fmt.Errorf("Couldn't parse argument, insert a valid number: %v", err)
			}
			limit = lmt
		default:
			limit = 2
	}

	pts, err := s.db.GetPostsForUser(context.Background(), int32(limit))

	if err != nil {
		return fmt.Errorf("Couldn't retrieve posts: %v", err)
	}

	for _, post := range pts {
		fmt.Printf("* %v\n", post.Title.String)
		fmt.Printf("  URL: %v\n", post.Url)
		if post.Description.Valid {
			fmt.Printf("  Description: %v\n", post.Description.String)
		}
		fmt.Printf("  Published: %v\n\n", post.PublishedAt.Time.Format("01 Jan 2006 15:04:05"))
	}

	return nil
}

