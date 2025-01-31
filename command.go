package main

import "fmt"

type command struct {
	name string
	args []string
}

type commands struct {
	mp map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.mp[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, bool := c.mp[cmd.name]
	if !bool {
		return fmt.Errorf("Command not found")
	}
	return f(s, cmd)
}
