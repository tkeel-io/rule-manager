package source

import "context"

type SourceConfig struct {
	Endpoint  string
	Name      string
	ProtoType string
}

type SourceTransport interface {
	Open(ctx context.Context) error
	Run(ctx context.Context) error
	Close(ctx context.Context) error
}

type Driver interface {
	NewSourceTransport(ctx context.Context, endpoint string, pt string) (SourceTransport, error)
}

type Configuration interface {
	Name() string
	GetString(key string) string
}
