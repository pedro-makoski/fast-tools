package handler

import (
	"github.com/pedro-makoski/fast-tools/comands"
	"github.com/pedro-makoski/fast-tools/comands/converter"
	"github.com/pedro-makoski/fast-tools/comands/password"
	"github.com/pedro-makoski/fast-tools/comands/sshconf"
)

var allComands = map[string]comands.IComand{
	"password": &password.PasswordGen{},
	"ssh":      &sshconf.SSHConf{},
	"convert":  &converter.ConverterConf{},
}
