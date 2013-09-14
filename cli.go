package main

import (
	"flag"
	"fmt"
	"log"
	"net/rpc"

	"github.com/arunjit/gourls/api"
)

// Flags
var (
	serverFlag = flag.String("server", "127.0.0.1:2001", "The RPC server")
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

	client, err := rpc.DialHTTPPath("tcp", *serverFlag, "/rpc")
	if err != nil {
		log.Fatalln("Error connecting to RPC server.", err)
	}

	var out string
	switch cmd.cmd {
	case cmdAdd:
		args := api.NewArgs(cmd.url)
		reply := new(api.NewReply)
		if err := client.Call("URL.New", args, reply); err != nil {
			log.Fatalf("Couldn't add %s\n%s\n", cmd.url, err.Error())
		} else {
			out = string(*reply)
		}
	case cmdSet:
		args := &api.SetArgs{Key: cmd.key, URL: cmd.url}
		reply := new(api.SetReply)
		if err := client.Call("URL.Set", args, reply); err != nil {
			log.Fatalf("Couldn't set %s => %s\n%s\n", cmd.key, cmd.url, err.Error())
		} else {
			out = cmd.key
		}
	case cmdGet:
		args := api.GetArgs(cmd.url)
		reply := new(api.GetReply)
		if err := client.Call("URL.Get", args, reply); err != nil {
			log.Fatalf("Couldn't get %s\n%s\n", cmd.key, err.Error())
		} else {
			out = string(*reply)
		}
	}

	fmt.Println(out)
}
