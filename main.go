package main

import (
	"fmt"
	"os/exec"

	"github.com/krosshykes/netonly/runas"
)

func main() {
	user := "test"
	logonDomain := "tDomain"
	pw := "test"
	path, _ := exec.LookPath("powershell")
	curDir := "C:\\users"

	err := runas.StartProcess(user, logonDomain, pw, path, curDir)
	if err != nil {
		fmt.Println(err)
	}
}
