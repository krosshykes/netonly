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
		err := runas.StartProcess(user, logonDomain, pw, nPath, curDir)
		if err != nil {
			fmt.Println(err)
		}
	}
	// user = "test"
	// logonDomain = "Domain"
	// pw = "test"

	// curDir = "C:\\users"
	// lenA := len(os.Args)
	// if lenA < 2 {
	// 	fmt.Println("Usage:")
	// 	fmt.Println("go run main.go username domain password appname directory")
	// } else if lenA < 3 {
	// 	switch os.Args[1] {
	// 	case "-h", "--help":
	// 		fmt.Println("Usage:")
	// 		fmt.Println("go run main.go username domain password appname directory")
	// 	default:
	// 		fmt.Println("Invalid Argument")
	// 	}
	// } else if lenA < 5 || lenA > 6 {
	// 	fmt.Println("Invalid number of Arguments")
	// } else {
	// 	user = os.Args[1]
	// 	logonDomain = os.Args[2]
	// 	pw = os.Args[3]
	// 	path, _ = exec.LookPath(os.Args[4])
	// 	if lenA == 6 {
	// 		curDir = os.Args[5]
	// 	}
	// 	err := runas.StartProcess(user, logonDomain, pw, path, curDir)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }

}
