package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var vari string = " "
	for vari != "y" && vari != "n" {
		fmt.Print("Did you installed git already?(y/n) ")
		fmt.Scan(&vari)
	}
	if vari == "n" {
		fmt.Println("Please install git before upgrading.(https://git-scm.com/downloads)")
		fmt.Scan(&vari)
		os.Exit(0)
	}

	fmt.Println("adding new files...")
	exec.Command("git", "add", ".").Run()

	fmt.Println("commiting current working directory...")
	exec.Command("git", "commit", "-m", "\"update\"").Run()

	fmt.Println("Please create a token here: https://github.com/settings/tokens account: it.rojdax.com@gmail.com")

	vari = " "
	for vari == " " {
		fmt.Print("Please enter the personal access token? ")
		fmt.Scan(&vari)
	}

	fmt.Println("pushing to github...")
	exec.Command("git", "push",
		"https://"+vari+"@github.com/Rj-22/rj-22.github.io.git").Run()
}
