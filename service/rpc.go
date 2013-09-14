package service

import (
	"strings"

	"github.com/arunjit/gourls/api"
)

type RPCService struct {
	store api.Store
}

func NewRPCService(s api.Store) *RPCService {
	return &RPCService{s}
}

func (s *RPCService) New(args *api.NewArgs, reply *api.NewReply) error {
	key, err := s.store.New(args.URL)
	if err != nil {
		return err
	}
	reply.Keys = strings.Split(key, ",")
	return nil
}

func (s *RPCService) Set(args *api.SetArgs, reply *api.SetReply) error {
	key, err := s.store.Set(args.Key, args.URL)
	if err != nil {
		return err
	}
	reply.Keys = strings.Split(key, ",")
	return nil
}

func (s *RPCService) Get(args *api.GetArgs, reply *api.GetReply) error {
	url, err := s.store.Get(args.Key)
	if err != nil {
		return err
	}
	reply.URL = url
	return nil
}
