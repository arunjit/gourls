package service

import (
	"github.com/arunjit/gourls/api"
)

type RPCService struct {
	store api.Store
}

func NewRPCService(s api.Store) *RPCService {
	return &RPCService{s}
}

func (s *RPCService) New(args *api.NewArgs, reply *api.NewReply) error {
	key, err := s.store.New(string(*args))
	if err != nil {
		return err
	}
	*reply = api.NewReply(key)
	return nil
}

func (s *RPCService) Set(args *api.SetArgs, reply *api.SetReply) error {
	if err := s.store.Set(args.Key, args.URL); err != nil {
		*reply = api.SetReply(false)
		return err
	}
	*reply = api.SetReply(true)
	return nil
}

func (s *RPCService) Get(args *api.GetArgs, reply *api.GetReply) error {
	url, err := s.store.Get(string(*args))
	if err != nil {
		return err
	}
	*reply = api.GetReply(url)
	return nil
}
