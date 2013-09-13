package main

import (
	"flag"
	"log"
	"net/http"
	"net/rpc"

	//"github.com/arunjit/gohttpx/httpx"
	"github.com/arunjit/gourls/service"
	"github.com/arunjit/gourls/store"
)

// Flags
var (
	rpcFlag   = flag.Bool("rpc", false, "Run HTTP RPC")
	jsonFlag  = flag.Bool("json", false, "Run JSON RPC")
	addrFlag  = flag.String("addr", "", "Address to run the server on.")
	redisFlag = flag.String("redis", "127.0.0.1:6379", "Redis host:port.")
)

const (
	serviceName = "URL"
	rpcPath     = "/rpc"
	debugPath   = "/_/debug/rpc"
)

func createRPC() {
	// The RPC service
	rpcSvc := service.NewRPCService(store.NewURLStore(*redisFlag))

	svr := rpc.NewServer()
	svr.RegisterName(serviceName, rpcSvc)
	svr.HandleHTTP(rpcPath, debugPath)

	rpc.Register(rpcSvc)
	rpc.HandleHTTP()
}

func createJSON() {
	log.Fatalln("Coming soon.")
}

func main() {
	flag.Parse()
	if *addrFlag == "" {
		log.Fatalln("-addr is required")
	}

	if !*rpcFlag && !*jsonFlag {
		log.Fatalln("At least one of -rpc or -json is required")
	}

	if *rpcFlag {
		createRPC()
	}
	if *jsonFlag {
		createJSON()
	}

	http.ListenAndServe(*addrFlag, nil)
}
