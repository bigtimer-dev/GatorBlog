package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/bigtimer-dev/GatorBlog/internal/database"
	"github.com/google/uuid"
)

// handler for login into database
func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("Handler Expect a Username in args")
	}

	user, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		fmt.Printf("User dont exist rigister a name first: %v \n", err)
		os.Exit(1)
	}

	err = s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		fmt.Printf("Error setting the user: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("User: %v have been set succesfully\n", user)

	return nil
}

// handler for register into database
func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("Handler expect a Username in args")
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	})
	if err != nil {
		fmt.Print("database or duplicate name error\n")
		os.Exit(1)
	}

	err = s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("User Created: %v\n", user)

	return nil
}
