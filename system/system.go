package system

import (
	"bytes"
	"os"
	"os/exec"
)

func Capture(command string, args []string) (string, string, error) {
	var outb, errb bytes.Buffer
	cmd := exec.Command(command, args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	err := cmd.Run()
	if err != nil {
		return "", "", err
	}
	return outb.String(), errb.String(), nil
}

func Nvim(args []string) error {
	cmd := exec.Command("nvim", args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
