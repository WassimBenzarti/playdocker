package lib

import (
	"fmt"
	"os"
	"os/exec"
)

func LaunchSSHTunnel(uri string) {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	if _, err = os.Stat(home + "/.ssh/id_rsa"); os.IsNotExist(err) {
		fmt.Println("Warning: It looks like you don't have any ssh keys configured. Run ssh-keygen command to initialize them.")
	}

	cmd := exec.Command("ssh", uri, "-i", home+"/.ssh/")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
