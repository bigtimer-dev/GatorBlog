package main

import "errors"

// struct of a command
type command struct {
	Name string
	Args []string
}

// struct of multiple commands
type commands struct {
	registeredCommands map[string]func(*state, command) error
}

// method to add a command to the commands struct
func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

// method to run a command registered in the commands structs
func (c *commands) run(s *state, cmd command) error {
	value, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return errors.New("command dont exist")
	}
	return value(s, cmd)
}
