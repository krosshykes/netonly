package main

import (
	"fmt"
	"os/exec"
)

func main() {
	user := "test"
	logonDomain := "tDomain"
	pw := "test"
	path, _ := exec.LookPath("cmd")
	curDir := "C:\\users"

	err := StartProcess(user, logonDomain, pw, path, curDir)
	if err != nil {
		fmt.Println(err)
	}
}
