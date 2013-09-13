package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/arunjit/gourls/store"
)

// Flags
var (
	redisFlag = flag.String("redis", "localhost:6379", "Redis server")
)

const (
	usage = `
	add <url>            adds the URL and generates a random key
	set <key> <url>      sets the URL to the given key (must be unique)
	get <key>            gets the URL for the given key
	`
)

// commands
const (
	cmdAdd = iota
	cmdSet
	cmdGet
)

type command struct {
	cmd      int
	key, url string
}

func printUsage() {
	log.Fatalf("Invalid command\n%s", usage)
}

func getCommand(args []string) *command {
	if len(args) < 2 {
		printUsage()
	}
	switch args[0] {
	case "add":
		return &command{cmdAdd, "", args[1]}
	case "set":
		return &command{cmdSet, args[1], args[2]}
	case "get":
		return &command{cmdGet, args[1], ""}
	default:
		printUsage()
	}
	// should never get here
	return nil
}

func main() {
	flag.Parse()
	cmd := getCommand(flag.Args())

	s := store.NewURLStore(*redisFlag)

	var out string
	switch cmd.cmd {
	case cmdAdd:
		if key, err := s.New(cmd.url); err != nil {
			log.Fatalf("Couldn't add %s\n%s\n", cmd.url, err.Error())
		} else {
			out = key
		}
	case cmdSet:
		if err := s.Set(cmd.key, cmd.url); err != nil {
			log.Fatalf("Couldn't set %s => %s\n%s\n", cmd.key, cmd.url, err.Error())
		} else {
			out = cmd.key
		}
	case cmdGet:
		if url, err := s.Get(cmd.key); err != nil {
			log.Fatalf("Couldn't get %s\n%s\n", cmd.key, err.Error())
		} else {
			out = url
		}
	}

	fmt.Println(out)
}
