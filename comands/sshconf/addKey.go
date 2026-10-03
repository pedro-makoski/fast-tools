package sshconf

import (
	"errors"
	"fmt"
)

func (s *SSHConf) addKey() error {
	if len(s.params) < 4 {
		return errors.New("Insuficient params length")
	}

	label := s.params[0]
	name := s.params[1]
	host := s.params[2]
	password := s.params[3]

	err := s.addKeyInFile(label, name, host, password)

	if err == nil {
		fmt.Println("Senha adicionada com sucesso")
	}

	return err
}

func (s *SSHConf) addKeyInFile(label string, name string, host string, password string) error {
	content, err := read()

	if err != nil {
		return err
	}

	content = append(content, SshStruct{
		Label:    label,
		Name:     name,
		Password: password,
		Host:     host,
	})

	err = save(content)

	return err
}
