package main

import (
	"fmt"
	"movie/cmd/database"
	"os"
	"slices"
)

var helpMessage = `Usage: go run cmd/database/script.go [options]
Options:
  -m, --migrate     Migrate tables
  -s, --seed        Insert data into tables
  -h, --help        Display this help message
`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(helpMessage)
		return
	}

	args := os.Args[1:]

	if slices.Contains(args, "-h") || slices.Contains(args, "--help") {
		fmt.Println(helpMessage)
		return
	}

	db := database.Connect()
	defer db.Close()

	actions := map[string]func(){
		"-m":        func() { database.Migrate(db) },
		"--migrate": func() { database.Migrate(db) },
		"-s":        func() { database.Seeds(db) },
		"--seed":    func() { database.Seeds(db) },
	}

	executed := false
	for _, arg := range args {
		if action, ok := actions[arg]; ok {
			action()
			executed = true
		}
	}

	if !executed {
		fmt.Println("[!] Unknown option.")
		fmt.Println(helpMessage)
	}
}
