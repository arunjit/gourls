package main

import (
	"flag"
	"log"
	"net/http"
	"net/rpc"

	grpc "github.com/gorilla/rpc/v2"
	jsrpc "github.com/gorilla/rpc/v2/json2"

	//"github.com/arunjit/gohttpx/httpx"
	"github.com/arunjit/gourls/api"
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
	serviceName  = "URL"
	rpcPath      = "/rpc"
	rpcDebugPath = "/_/debug/rpc"
	jsonPath     = "/json"
)

func newStore() api.Store {
	return store.NewURLStore(*redisFlag)
}

func newService() api.Service {
	return service.NewRPCService(newStore())
}

func createRPC() {
	// The RPC service
	svc := newService()

	svr := rpc.NewServer()
	svr.RegisterName(serviceName, svc)
	svr.HandleHTTP(rpcPath, rpcDebugPath)
	svr.Register(svc)
	svr.HandleHTTP()
}

func createJSON() {
	svc := service.NewJSONService(newService())

	svr := grpc.NewServer()
	svr.RegisterCodec(jsrpc.NewCodec(), "application/json")
	svr.RegisterService(svc, serviceName)
	http.Handle(jsonPath, svr)
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
