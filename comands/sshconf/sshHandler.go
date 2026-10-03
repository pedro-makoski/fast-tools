package sshconf

import (
	"errors"
	"os"
)

type SSHConf struct {
	params []string
}

func (s *SSHConf) Do() error {
	return s.handleComand()
}

func (s *SSHConf) handleComand() error {
	if len(os.Args) <= 2 {
		return errors.New("It's required more arguments")
	}

	comand := os.Args[2]
	s.params = os.Args[3:]

	switch comand {
	case "add":
		return s.addKey()
	case "enter":
		return s.enterSsh()
	}

	return errors.New("Don't exist this comand: " + comand)
}
