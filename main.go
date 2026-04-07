package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/bigtimer-dev/GatorBlog/internal/config"
	"github.com/bigtimer-dev/GatorBlog/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	// read the config in the .gatorjson
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	// Openening the database
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatal(err)
	}

	// Creating a new data base
	dbQueries := database.New(db)

	// new instance of state initiallized with the read config file and the queries from database
	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}
	// new instance of commands initiallized
	newCommands := &commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	// register our commands
	newCommands.register("login", handlerLogin)
	newCommands.register("register", handlerRegister)
	newCommands.register("reset", handlerReset)
	newCommands.register("users", handlerList)

	// verify if the cli line contain command name and arguments
	if len(os.Args) < 2 {
		fmt.Printf("Not enough arguments\n")
		os.Exit(1)
	}
	// take argument 1 as name of command and arg 2 as the argument for the command
	newCommand := command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}
	// running the command with the arguments passed
	err = newCommands.run(programState, newCommand)
	if err != nil {
		log.Fatal(err)
	}
}
