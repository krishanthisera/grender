package backend

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (b S3) Put(url string, data []byte) error {
	_, err := b.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &b.BucketName,
		Key:    &url,
		Body:   bytes.NewReader(data),
	})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (b S3) Get(url string) ([]byte, error) {
	res, err := b.S3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &b.BucketName,
		Key:    &url,
	})

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
