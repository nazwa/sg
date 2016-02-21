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
	return len(output) == 0 && err == nil
}

func main() {
	cwd, _ := os.Getwd()
	if !isInGit(cwd) {
		fmt.Println("This is not a git repo :(")
		return
	}

	message := "Updates..."
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

	pullCmd := exec.Command("git", "pull")
	pullCmd.Dir = cwd
	if err := pullCmd.Run(); err != nil {
		fmt.Println(err.Error())
		return
	}

	pushCmd := exec.Command("git push")
	pushCmd.Dir = cwd
	if err := pushCmd.Run(); err != nil {
		fmt.Println(err.Error())
		return
	}
}
