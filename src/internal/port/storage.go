package port

import "context"

type StorageClient interface {
	Upload(ctx context.Context, bucket, key string, data []byte, contentType string) (string, error)
	GetURL(ctx context.Context, bucket, key string) (string, error)
	Delete(ctx context.Context, bucket, key string) error
	BucketExists(ctx context.Context, bucket string) (bool, error)
	CreateBucket(ctx context.Context, bucket string) error
}
