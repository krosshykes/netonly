package main

import (
	"fmt"
	"os"
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
	user = "test"
	logonDomain = "Domain"
	pw = "test"

	curDir = "C:\\users"
	lenA := len(os.Args)
	if lenA < 2 {
		fmt.Println("Usage:")
		fmt.Println("go run main.go username domain password appname directory")
	} else if lenA < 3 {
		switch os.Args[1] {
		case "-h", "--help":
			fmt.Println("Usage:")
			fmt.Println("go run main.go username domain password appname directory")
		default:
			fmt.Println("Invalid Argument")
		}
	} else if lenA < 5 {
		fmt.Println("Invalid number of Arguments")
	} else {
		user = os.Args[1]
		logonDomain = os.Args[2]
		pw = os.Args[3]
		path, _ = exec.LookPath(os.Args[4])
		if lenA == 6 {
			curDir = os.Args[5]
		}
		err := runas.StartProcess(user, logonDomain, pw, path, curDir)
		if err != nil {
			fmt.Println(err)
		}
	}

}
