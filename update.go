package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Please install git before upgrading.(https://git-scm.com/downloads)")
	var git_installed byte
	fmt.Scanf("%c", &git_installed)
	out, err := exec.Command("git", "add", ".").Output()
	out, err = exec.Command("git", "commit", "-m", "update website").Output()
	out, err = exec.Command("git", "push", "-u", "https://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@github.com/rj-22/rj-22.github.io.git").Output()
	out, err = exec.Command("git", "pull", "https://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@github.com/rj-22/rj-22.github.io.git").Output()
	fmt.Println(out)
	fmt.Println(err)
}
