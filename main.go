package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/abtsousa/gator/internal/config"
	"github.com/abtsousa/gator/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {

	return func(s *state, cmd command) error {
		
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("Couldn't retrieve current user: %v", err)
		}
		err = handler(s, cmd, user)
		if err != nil {
			return err
		}
		return nil
	}
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	dbURL := cfg.DbURL
	db, err := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)

	s := state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := commands{mp: make(map[string]func(*state, command) error)}
	cmds.register("login", handler_login)
	cmds.register("register", handler_register)
	cmds.register("users", handler_users)
	cmds.register("agg", handler_agg)
	cmds.register("addfeed", middlewareLoggedIn(handler_addfeed))
	cmds.register("feeds", handler_feeds)
	cmds.register("follow", middlewareLoggedIn(handler_follow))
	cmds.register("following", middlewareLoggedIn(handler_following))
	cmds.register("unfollow", middlewareLoggedIn(handler_unfollow))
	cmds.register("reset", handler_reset)
	cmd := command{}

	if len(os.Args) < 2 {
		log.Fatalf("You must provide a command and an argument.")
		return
	}
	cmd.name, cmd.args = os.Args[1], os.Args[2:]
	err = cmds.run(&s, cmd)
	if err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
