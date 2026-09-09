package storage

import (
    "context"
    "fmt"
    "time"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/credentials"
    "github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
    Client *s3.Client
    Bucket string
    UsePathStyle bool
}

func NewS3(endpoint, accessKey, secretKey, bucket string, pathStyle bool) (*S3Client, error) {
    cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion("auto"), config.WithCredentialsProvider(credentials.StaticCredentialsProvider{Value: aws.Credentials{AccessKeyID: accessKey, SecretAccessKey: secretKey}}))
    if err != nil { return nil, err }
    client := s3.NewFromConfig(cfg, func(o *s3.Options) {
        o.BaseEndpoint = aws.String(endpoint)
        o.UsePathStyle = pathStyle
    })
    return &S3Client{Client: client, Bucket: bucket, UsePathStyle: pathStyle}, nil
}

func (s *S3Client) PresignedPut(key string, contentType string, ttl time.Duration) (string, error) {
    // Basit örnek — gerçek kullanımda presign client gerekir.
    return fmt.Sprintf("presigned-url-for-%s", key), nil
}
