package storage

import (
	"context"
	"errors"
	"fmt"
)

type ShortedURL string
type URLKey string

var (
	StorageError = errors.New("storage")
	ErrCollision = fmt.Errorf("%w.collision", StorageError)
	ErrNotFound  = fmt.Errorf("%w.not_found", StorageError)
)

type Storage interface {
	PutURL(ctx context.Context, url ShortedURL) (URLKey, error)
	GetURL(ctx context.Context, key URLKey) (ShortedURL, error)
}
