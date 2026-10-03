package handler

import (
	"github.com/pedro-makoski/fast-tools/comands"
	"github.com/pedro-makoski/fast-tools/comands/password"
)

var allComands = map[string]comands.IComand{
	"password": &password.PasswordGen{},
	"convert":  nil,
	"ssh":      nil,
}
