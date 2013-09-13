package service

import (
	"github.com/arunjit/gourls/api"
)

type JSONService struct {
	rpc api.Service
}

func NewJSONService(s api.Store) *JSONService {
	return &JSONService{NewRPCService(s)}
}
