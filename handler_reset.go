package main

import (
	"context"
	"fmt"
)

func handler_reset(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Usage: %v", cmd.name)
	}
	err := s.db.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("Error deleting user database: %v", err)
	}
	fmt.Println("Database reset successfully!")
	return nil
}
