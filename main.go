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

func runCmd(wd, app string, args ...string) bool {
	cmd := exec.Command(app, args...)
	cmd.Dir = wd

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("=============================")
		fmt.Println("Error :( git had this to say:")
		fmt.Println("=============================")
		fmt.Printf("%s", output)
		fmt.Println("=============================")
		return false
	}
	return true
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

	if !runCmd(cwd, "git", "add", "-A") {
		return
	}
	if !runCmd(cwd, "git", "commit", "-am\""+message) {
		return
	}
	if !runCmd(cwd, "git", "pull") {
		return
	}
	if !runCmd(cwd, "git", "push") {
		return
	}

}
