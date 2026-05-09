package storage

import (
	"bytes"
	"context"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"sai-server/internal/port"
)

type MinIO struct {
	client *minio.Client
}

func NewMinIO(endpoint, accessKey, secretKey string) (*MinIO, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}
	return &MinIO{client: client}, nil
}

func NewMinIOOrNoop(endpoint, accessKey, secretKey string) port.StorageClient {
	c, err := NewMinIO(endpoint, accessKey, secretKey)
	if err != nil {
		return &noopStorage{}
	}
	return c
}

func (m *MinIO) Upload(ctx context.Context, bucket, key string, data []byte, contentType string) (string, error) {
	exists, err := m.client.BucketExists(ctx, bucket)
	if err != nil {
		return "", err
	}
	if !exists {
		if err := m.client.MakeBucket(ctx, bucket, minio.MakeBucketOptions{}); err != nil {
			return "", err
		}
	}

	reader := bytes.NewReader(data)
	_, err = m.client.PutObject(ctx, bucket, key, reader, int64(len(data)), minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return "", err
	}

	return key, nil
}

func (m *MinIO) GetURL(ctx context.Context, bucket, key string) (string, error) {
	u, err := m.client.PresignedGetObject(ctx, bucket, key, 7*24*time.Hour, nil)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

func (m *MinIO) Delete(ctx context.Context, bucket, key string) error {
	return m.client.RemoveObject(ctx, bucket, key, minio.RemoveObjectOptions{})
}

func (m *MinIO) BucketExists(ctx context.Context, bucket string) (bool, error) {
	return m.client.BucketExists(ctx, bucket)
}

func (m *MinIO) CreateBucket(ctx context.Context, bucket string) error {
	return m.client.MakeBucket(ctx, bucket, minio.MakeBucketOptions{})
}

type noopStorage struct{}

func (n *noopStorage) Upload(_ context.Context, _, key string, _ []byte, _ string) (string, error) {
	return key, nil
}

func (n *noopStorage) GetURL(_ context.Context, _, key string) (string, error) {
	return key, nil
}

func (n *noopStorage) Delete(_ context.Context, _, _ string) error {
	return nil
}

func (n *noopStorage) BucketExists(_ context.Context, _ string) (bool, error) {
	return true, nil
}

func (n *noopStorage) CreateBucket(_ context.Context, _ string) error {
	return nil
}
