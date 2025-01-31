package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/abtsousa/gator/internal/database"
	"github.com/google/uuid"
)

func handler_register(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Usage: %v <username>", cmd.name)
	}

	ctx := context.Background()
	if _, err := s.db.GetUser(ctx, cmd.args[0]); err == nil {
		log.Fatal("A user with that name already exists.")
		os.Exit(1)
	}
	user, err := s.db.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.args[0],
	})

	if err != nil {
		return fmt.Errorf("An error occurred creating the user: %v", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("An error occurred setting the current user: %v", err)
	}

	fmt.Println("The user was successfully created.")
	log.Printf("Created user %s", user.Name)

	return nil

}
