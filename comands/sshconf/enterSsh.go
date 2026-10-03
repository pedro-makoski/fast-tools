package sshconf

import (
	"errors"

	"github.com/pedro-makoski/fast-tools/utils"
)

func (s *SSHConf) enterSsh() error {
	if len(s.params) == 0 {
		return errors.New("Insuficient params length")
	}

	return s.enterOnLabel(s.params[0])
}

func (s *SSHConf) enterOnLabel(label string) error {
	listOfKeys, err := s.getKeys()

	if err != nil {
		return err
	}

	keyConf, err := s.searchOnLabel(label, listOfKeys)

	if err != nil {
		return err
	}

	return keyConf.login()
}

func (s *SSHConf) getKeys() ([]SshStruct, error) {
	content, err := read()

	if err != nil {
		return []SshStruct{}, err
	}

	keys := utils.Revert(content)

	return keys, nil
}

func (s *SSHConf) searchOnLabel(label string, items []SshStruct) (SshStruct, error) {
	for _, v := range items {
		if v.Label == label {
			return v, nil
		}
	}

	return SshStruct{}, errors.New("No acess with this mode")
}
