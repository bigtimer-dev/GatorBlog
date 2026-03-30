package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("Handler Expect a Username in args")
	}
	err := s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}
	fmt.Print("The username have been set succesfully")
	return nil
}
