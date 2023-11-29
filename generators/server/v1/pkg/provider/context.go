package provider

import (
	"context"
	"fmt"
)

type Context struct {
	Region      string
	TypeName    string
	AccessToken string
}

const (
	Key = "iacker-context"
)

func WithContext(ctx context.Context, credential *Context) context.Context {
	return context.WithValue(ctx, Key, credential)
}

func FromContext(ctx context.Context) (*Context, error) {
	v := ctx.Value(Key)

	cred, ok := v.(*Context)
	if !ok {
		return nil, fmt.Errorf("no credential found")
	}
	return cred, nil
}
