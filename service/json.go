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

func (s *JSONService) New(r *http.Request, args *api.NewArgs, result *api.NewResult) error {
	return s.rpc.New(args, result)
}

func (s *JSONService) Set(r *http.Request, args *api.SetArgs, result *api.SetResult) error {
	return s.rpc.Set(args, result)
}

func (s *JSONService) Get(r *http.Request, args *api.GetArgs, result *api.GetResult) error {
	return s.rpc.Get(args, result)
}
