package main

import (
	"flag"
	"fmt"
	"os/exec"

	"github.com/krosshykes/netonly/runas"
)

func main() {
	var (
		user        string
		logonDomain string
		pw          string
		path        string
		curDir      string
	)

	flag.StringVar(&user, "u", "test", "Username")
	flag.StringVar(&logonDomain, "d", "Domain", "Domain")
	flag.StringVar(&pw, "p", "pass", "Password")
	flag.StringVar(&path, "a", "powershell", "Application Name")
	flag.StringVar(&curDir, "cd", "C:\\users", "Current Directory")
	flag.Parse()
	nPath, pErr := exec.LookPath(path)
	if pErr != nil {
		fmt.Println("Invalid Application Name: ", pErr)
	} else {
		err := runas.StartProcess1(user, logonDomain, pw, nPath, curDir)
		if err != nil {
			fmt.Println(err)
		}
	}
}
