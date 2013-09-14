package service

import (
	"net/http"

	"github.com/arunjit/gourls/api"
)

// JSONService wraps a bare RPC service and serves it as JSON-RPC.
type JSONService struct {
	rpc api.Service
}

func NewJSONService(s api.Service) *JSONService {
	return &JSONService{s}
}

func (s *JSONService) New(r *http.Request, args *api.NewArgs, reply *api.NewReply) error {
	return s.rpc.New(args, reply)
}

func (s *JSONService) Set(r *http.Request, args *api.SetArgs, reply *api.SetReply) error {
	return s.rpc.Set(args, reply)
}

func (s *JSONService) Get(r *http.Request, args *api.GetArgs, reply *api.GetReply) error {
	return s.rpc.Get(args, reply)
}
