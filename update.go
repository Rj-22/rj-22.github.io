package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("git", "add", ".").Output()
	fmt.Println(out)
	fmt.Println(err)
	out, err = exec.Command("git", "commit", "-m", "update website").Output()
	fmt.Println(out)
	fmt.Println(err)
	out, err = exec.Command("git", "push", "-u", "https://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@github.com/rj-22/rj-22.github.io.git").Output()
	fmt.Println(out)
	fmt.Println(err)
	out, err = exec.Command("git", "pull", "https://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@github.com/rj-22/rj-22.github.io.git").Output()
	fmt.Println(out)
	fmt.Println(err)
}
