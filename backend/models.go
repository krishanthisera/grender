package backend

import "github.com/aws/aws-sdk-go-v2/service/s3"

type Backend interface {
	Put(url string, data []byte) error
	Get(url string) ([]byte, error)
}

type S3 struct {
	BucketName string     `yaml:"bucketName"`
	Region     string     `yaml:"region"`
	S3Client   *s3.Client `yaml:"-"`
}

type FileSystem struct {
	BaseDir string `yaml:"baseDir"`
}
