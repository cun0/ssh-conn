package main

import (
	"os"
	"os/exec"
)

type SshConnectionOp struct {
	IpAdress string
	Password string
}

func (op SshConnectionOp) Run() error {
	cmd := exec.Command("ssh", "root@165.227.138.161")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	return err
}
