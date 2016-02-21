package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func isInGit(wd string) bool {
	cmd := exec.Command("git", "rev-parse")
	cmd.Dir = wd
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	return len(output) == 0
}

func main() {
	cwd, _ := os.Getwd()
	fmt.Println(isInGit(cwd))

	message := ""

	if len(os.Args) > 1 {
		words := os.Args[1:]
		message = strings.Join(words, " ")
	}
	addCmd := exec.Command("git", "add", "-A")
	addCmd.Dir = cwd

	if err := addCmd.Run(); err != nil {
		fmt.Println(err.Error())
		return
	}

	commitCmd := exec.Command("git", "commit", "-am\""+message)
	commitCmd.Dir = cwd

	if err := commitCmd.Run(); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(cwd)
	fmt.Println(message)
}
