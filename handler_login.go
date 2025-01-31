package main

import (
	"context"
	"fmt"
	"log"
	"os"
)

func handler_login(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Usage: %v <username>", cmd.name)
	}

	name := cmd.args[0]

	if _, err := s.db.GetUser(context.Background(), name); err != nil {
		log.Fatal("User not found!")
		os.Exit(1)
	}

	err := s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("Couldn't set current user: %v", err)
	}
	fmt.Printf("User %s was successfully set.", name)
	return nil
}
