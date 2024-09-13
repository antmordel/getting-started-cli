package main

import (
	"log"

	"github.com/antmordel/getting-started-cli/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("Whoops. There was an error executing the amazing getting-started-cli: %v", err)
	}
}
