package server

import (
	"github.com/arunjit/gourls/api"
)

type Store interface {
	New(url string) (string, error)
	Set(key, url string) error
	Get(key string) (string, error)
}

type urlService struct {
	store *Store
}

func NewRpcService(s *Store) *urlService {
	return &urlService{s}
}

func (s *urlService) New(args *api.NewArgs, result *api.NewResult) error {
	key, err := s.store.New(string(args))
	if err != nil {
		return err
	}
	*result = api.NewResult(key)
	return nil
}

func (s *urlService) Set(args *api.SetArgs, result *api.SetResult) error {
	if err := s.store.Set(args.Key, args.URL); err != nil {
		*result = api.SetResult(false)
		return err
	}
	*result = api.SetResult(true)
	return nil
}

func (s *urlService) Get(args *api.GetArgs, result *api.GetResult) error {
	url, err := s.store.Get(string(args))
	if err != nil {
		return err
	}
	*result = api.GetResult(url)
	return nil
}
