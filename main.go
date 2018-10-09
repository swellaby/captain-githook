package main

import (
	"fmt"
	"github.com/swellaby/captain-githook/internal/cli"
	"os"
)

func main() {
	if err := cli.GetRunner().Execute(); err != nil {
		fmt.Println("captain-githook encountered fatal error")
		os.Exit(1)
	}
}
