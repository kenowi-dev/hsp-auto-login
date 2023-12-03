package main

import (
	"fmt"
	"github.com/kenowi-dev/hsp-auto-login/cmd/httpushover/cmd"
	"os"
)

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
