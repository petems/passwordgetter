package main

import (
	"fmt"
	"os"

	"github.com/petems/passwordgetter/cmd"
	"github.com/sirupsen/logrus"
)

func main() {
	var pwdr cmd.StdInPasswordReader
	fmt.Println("Enter password: ")
	result, err := cmd.Run(pwdr)
	if err != nil {
		logrus.Fatalf("Something went wrong: %v", err)
		os.Exit(1)
	}

	fmt.Printf("\nYour password was: %s\n", result)
}
