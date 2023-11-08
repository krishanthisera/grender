package backend

import "fmt"

type S3 struct {
	BucketName string
}

func (s S3) Put(url string, data []byte) error {
	fmt.Println(string(data))
	return nil
}

func (s S3) Get(url string) ([]byte, error) {
	fmt.Printf("Getting %s to S3", s.BucketName)
	return nil, nil
}
