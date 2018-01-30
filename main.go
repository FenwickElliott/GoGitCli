package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fail()
	} else {
		switch strings.ToLower(args[0]) {
		case "help":
			help()
		default:
			fail()
		}
	}

}

func fail() {
	fmt.Println("Not a valid options. Use gogitcli help for a list of commands")
}

func help() {
	fmt.Println("gogitcli new [repo name]")
}
