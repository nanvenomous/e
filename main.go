package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func nvim(args []string) error {
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

func main() {
	fileName := "Session.vim"
	path := "./"
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	ndirs := strings.Split(wd, "/")
	for range ndirs {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			err := nvim([]string{})
			if err != nil {
				panic(err)
			}
			break
		} else {
			fullFilePath := path + fileName
			fmt.Println(fullFilePath)
			if _, err := os.Stat(fullFilePath); !os.IsNotExist(err) {
				err := nvim([]string{"-S", fullFilePath})
				if err != nil {
					panic(err)
				}
				break
			}
		}
		if path == "./" {
			path = "../"
		} else {
			path += "../"
		}
	}

}
