package password

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type PasswordGen struct {
	actual      string
	length      int
	hasSimbols  bool
	hasChars    bool
	hasNumbers  bool
	isMinusculo bool
	isMaiusculo bool
}

func (p *PasswordGen) Do() error {
	p.prepareComands(os.Args[2:])

	err := p.verifyParams()

	if err != nil {
		return err
	}

	senha, err := p.generatePassword()

	if err != nil {
		return err
	}

	fmt.Println("Password Generated with sucess:")
	fmt.Println(senha)

	return nil
}

func (p *PasswordGen) verifyParams() error {
	if !p.hasChars && !p.hasNumbers && !p.hasSimbols {
		return errors.New("You can't disable all type of chars")
	}

	return nil
}

func (p *PasswordGen) prepareComands(args []string) {
	passwordFlags := flag.NewFlagSet("password", flag.ExitOnError)

	length := passwordFlags.Int("size", 8, "Size of your password")

	capitalize := passwordFlags.String(
		"capitalize",
		"both",
		"Both for Upper and Lower case, up to upper case and low to lower case",
	)

	noSimbols := passwordFlags.Bool(
		"noSimbols",
		false,
		"Must don't exist simbols in your password",
	)

	noNumbers := passwordFlags.Bool(
		"noNumbers",
		false,
		"Must don't exist numbers in your password",
	)

	noLetters := passwordFlags.Bool(
		"noLetters",
		false,
		"Must don't exist letters in your password",
	)

	passwordFlags.Parse(args)

	p.length = *length
	p.hasSimbols = !*noSimbols
	p.hasNumbers = !*noNumbers
	p.hasChars = !*noLetters

	p.handleCapitalization(*capitalize)
}

func (p *PasswordGen) handleCapitalization(captialize string) error {
	switch captialize {
	case "both":
		p.isMaiusculo = true
		p.isMinusculo = true
	case "up":
		p.isMaiusculo = true
		p.isMinusculo = false
	case "low":
		p.isMaiusculo = false
		p.isMinusculo = true
	default:
		return errors.New("O tipo específicado para capitalização não existe")
	}

	return nil
}
