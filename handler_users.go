package main

import (
	"context"
	"fmt"
)

func handler_users(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	for _, user := range users {
		var isCurrent string
		if user.Name == s.cfg.CurrentUserName {
			isCurrent = " (current)"
		} else {
			isCurrent = ""
		}
		fmt.Printf("* %s%s\n", user.Name, isCurrent)
	}
	return err
}
