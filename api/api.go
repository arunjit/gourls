package api

type NewArgs string   // key
type NewResult string // URL

type SetArgs struct {
	Key, URL string
}
type SetResult bool

type GetArgs string   // key
type GetResult string // URL

type Service interface {
	New(*NewArgs, *NewResult) error
	Set(*SetArgs, *SetResult) error
	Get(*GetArgs, *GetResult) error
}

type Store interface {
	New(string) (string, error)
	Set(string, string) error
	Get(string) (string, error)
}
