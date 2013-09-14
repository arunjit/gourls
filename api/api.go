package api

type NewArgs string  // key
type NewReply string // URL

type SetArgs struct {
	Key, URL string
}
type SetReply bool

type GetArgs string  // key
type GetReply string // URL

type Service interface {
	New(*NewArgs, *NewReply) error
	Set(*SetArgs, *SetReply) error
	Get(*GetArgs, *GetReply) error
}

type Store interface {
	New(string) (string, error)
	Set(string, string) error
	Get(string) (string, error)
}
