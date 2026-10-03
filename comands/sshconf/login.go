package sshconf

import (
	"errors"
	"fmt"
	"net"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

func (s *SshStruct) login() error {
	config := &ssh.ClientConfig{
		User: s.Name,
		Auth: []ssh.AuthMethod{
			ssh.Password(s.Password),
			ssh.KeyboardInteractive(func(user, instruction string, questions []string, echos []bool) ([]string, error) {
				answers := make([]string, len(questions))
				for i := range answers {
					answers[i] = s.Password
				}
				return answers, nil
			}),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	}

	addr := net.JoinHostPort(s.Host, "22")
	fmt.Println("Conectando em", addr)

	conn, err := net.DialTimeout("tcp", addr, config.Timeout)
	if err != nil {
		return fmt.Errorf("erro TCP: %w", err)
	}
	conn.SetDeadline(time.Now().Add(15 * time.Second))

	sshConn, chans, reqs, err := ssh.NewClientConn(conn, addr, config)
	if err != nil {
		conn.Close()
		return fmt.Errorf("erro no handshake/autenticação SSH: %w", err)
	}
	conn.SetDeadline(time.Time{})

	client := ssh.NewClient(sshConn, chans, reqs)
	defer client.Close()

	fmt.Println("Conectado!")

	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("erro ao criar sessão: %w", err)
	}
	defer session.Close()

	fd := int(os.Stdin.Fd())

	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return fmt.Errorf("erro ao colocar terminal em modo raw: %w", err)
	}
	defer term.Restore(fd, oldState)

	width, height, err := term.GetSize(fd)
	if err != nil {
		return fmt.Errorf("erro ao obter tamanho do terminal: %w", err)
	}

	termType := os.Getenv("TERM")
	if termType == "" {
		termType = "xterm-256color"
	}

	if err := session.RequestPty(termType, height, width, ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}); err != nil {
		return fmt.Errorf("erro ao criar PTY: %w", err)
	}

	session.Stdin = os.Stdin
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	if err := session.Shell(); err != nil {
		return fmt.Errorf("erro ao iniciar shell: %w", err)
	}

	if err := session.Wait(); err != nil {
		var exitErr *ssh.ExitError
		if errors.As(err, &exitErr) {
			return nil
		}
		return err
	}
	return nil
}
