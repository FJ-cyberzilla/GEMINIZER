// cmd/cli/main.go
package main

import (
	"fmt"
	"os"
	
	"geminizer-enterprise/internal/core/ai"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: geminizer <command> [options]")
		fmt.Println("Commands: generate, history, admin, version")
		os.Exit(1)
	}
	
	command := os.Args[1]
	
	switch command {
	case "generate":
		handleGenerate()
	case "history":
		handleHistory()
	case "admin":
		handleAdmin()
	case "version":
		fmt.Println("Geminizer Enterprise v1.0.0")
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}
