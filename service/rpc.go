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

func (s *RPCService) New(args *api.NewArgs, result *api.NewResult) error {
	key, err := s.store.New(string(*args))
	if err != nil {
		return err
	}
	*result = api.NewResult(key)
	return nil
}

func (s *RPCService) Set(args *api.SetArgs, result *api.SetResult) error {
	if err := s.store.Set(args.Key, args.URL); err != nil {
		*result = api.SetResult(false)
		return err
	}
	*result = api.SetResult(true)
	return nil
}

func (s *RPCService) Get(args *api.GetArgs, result *api.GetResult) error {
	url, err := s.store.Get(string(*args))
	if err != nil {
		return err
	}
	*result = api.GetResult(url)
	return nil
}
