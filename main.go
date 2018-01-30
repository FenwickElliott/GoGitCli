package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fail()
	} else {
		switch strings.ToLower(args[0]) {
		case "new":
			new(args[1])
		case "help":
			help()
		default:
			fail()
		}
	}
}

func new(name string) {
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		log.Fatal("No access token present")
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: string(token)})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	repo := &github.Repository{
		Name: github.String(name),
	}
	_, _, err := client.Repositories.Create(ctx, "", repo)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("successfully created", name)
	}
}

func fail() {
	fmt.Println("Not a valid options. Use gogitcli help for a list of commands")
}

func help() {
	fmt.Println("gogitcli new [repo name]")
}
