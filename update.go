package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var git_installed string = " "
	for git_installed != "y" && git_installed != "n" {
		fmt.Print("Did you installed git already?(y/n) ")
		fmt.Scan(&git_installed)
	}
	if git_installed == "n" {
		fmt.Println("Please install git before upgrading.(https://git-scm.com/downloads)")
		fmt.Scan(&git_installed)
		os.Exit(0)
	}
	fmt.Println("adding new files...")
	out, err := exec.Command("git", "add", ".").Output()
	fmt.Println("commiting current working directory...")
	out, err = exec.Command("git", "commit", "-m", "update website").Output()
	fmt.Println("pushing to github...")
	out, err = exec.Command("git", "push", "-u", "https://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@github.com/Rj-22/rj-22.github.io.git").Output()
	if err != nil {
		fmt.Println(out)
		fmt.Println(err.Error())
	}
}
