package main

import (
	"fmt"
	"gogo/cmd/server/cmd"
	_ "net/http/pprof"
	"os"
)

func main() {
	fmt.Println("Hello World")
	if err := cmd.RootCmd.Execute(); err != nil {
		_, err2 := fmt.Fprintln(os.Stderr, err)
		if err2 != nil {
			return
		}
		os.Exit(1)
	}
}
