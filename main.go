package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bigtimer-dev/GatorBlog/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	// read the config in the .gatorjson
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	// new instance of state initiallized with the read config file
	programState := &state{
		cfg: &cfg,
	}
	// new instance of commands initiallized
	newCommands := &commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	// register our first commands the login one
	newCommands.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Printf("Not enough arguments\n")
		os.Exit(1)
	}

	newCommand := command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err = newCommands.run(programState, newCommand)
	if err != nil {
		log.Fatal(err)
	}
}
