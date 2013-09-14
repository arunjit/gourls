package api

type NewArgs struct {
	URL string `json:"url"`
}
type NewReply struct {
	Keys []string `json:"keys"`
}

type SetArgs struct {
	Key string `json:"key"`
	URL string `json:"url"`
}
type SetReply struct {
	Keys []string `json:"keys"`
}

type GetArgs struct {
	Key string `json:"key"`
}
type GetReply struct {
	URL string `json:"url"`
}

type Service interface {
	New(*NewArgs, *NewReply) error
	Set(*SetArgs, *SetReply) error
	Get(*GetArgs, *GetReply) error
}

type Store interface {
	New(string) (string, error)
	Set(string, string) (string, error)
	Get(string) (string, error)
}
