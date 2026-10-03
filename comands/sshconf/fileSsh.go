package sshconf

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type SshStruct struct {
	Label    string `json:"label"`
	Host     string `json:"host"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

var nameJson = "password.json"

func getPath() (string, error) {
	exe, err := os.Executable()

	if err != nil {
		return "", err
	}

	exeDir := filepath.Dir(exe)
	filePath := filepath.Join(exeDir, nameJson)

	return filePath, nil
}

func save(items []SshStruct) error {
	data, err := json.MarshalIndent(items, "", "    ")

	if err != nil {
		return err
	}

	path, err := getPath()

	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func read() ([]SshStruct, error) {
	path, err := getPath()

	if err != nil {
		return []SshStruct{}, err
	}

	_, err = os.Stat(path)

	var items = []SshStruct{}

	if err == nil {
		content, err := os.ReadFile(path)

		if err != nil {
			return []SshStruct{}, err
		}

		err = json.Unmarshal(content, &items)

		if err != nil {
			return []SshStruct{}, err
		}
	}

	return items, nil
}
