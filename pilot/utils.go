package pilot

import (
	"context"
	"fmt"
	"os"

	aws "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/krishanthisera/grender/backend"
	"gopkg.in/yaml.v2"
)

func createBackendFromConfig(backendConfig interface{}) (backend.Backend, error) {
	switch b := backendConfig.(type) {
	case backend.S3:
		cfg, err := aws.LoadDefaultConfig(context.TODO(), aws.WithRegion(fmt.Sprintf(b.Region)))
		bucket := backend.S3{BucketName: "grender.io", S3Client: s3.NewFromConfig(cfg)}
		return &bucket, err
	case backend.FileSystem:
		fs := backend.FileSystem{BaseDir: b.BaseDir}
		return &fs, nil
	default:

		panic(fmt.Sprintf("Unknown backend: %T", b))
	}
}

// TO DO: Refactor this to use the backend package
// func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
// 	var buf map[string]interface{}
// 	err := unmarshal(&buf)
// 	if err != nil {
// 		return err
// 	}

// 	c.Version = buf["version"].(string)

// 	renderingConfig, ok := buf["renderingConfig"].(map[interface{}]interface{})
// 	if !ok {
// 		return errors.New("renderingConfig field not found or is not a map")
// 	}

// 	renderingConfigData, err := yaml.Marshal(renderingConfig)
// 	if err != nil {
// 		return err
// 	}
// 	err = yaml.Unmarshal(renderingConfigData, &c.RenderingConfig)
// 	if err != nil {
// 		return err
// 	}

// 	server, ok := buf["server"].(map[interface{}]interface{})
// 	if !ok {
// 		return errors.New("server field not found or is not a map")
// 	}

// 	c.Server.Port, ok = server["port"].(string)
// 	if !ok {
// 		return errors.New("port field not found or is not a string")
// 	}

// 	backend, ok := buf["backend"].(map[interface{}]interface{})
// 	if !ok {
// 		return errors.New("backend field not found or is not a map")
// 	}

// 	if s3, ok := backend["s3"].(map[interface{}]interface{}); ok {
// 		s3Data, err := yaml.Marshal(s3)
// 		if err != nil {
// 			return err
// 		}
// 		err = yaml.Unmarshal(s3Data, &c.Backend.S3)
// 		if err != nil {
// 			return err
// 		}

// 	} else if fs, ok := backend["fileSystem"].(map[interface{}]interface{}); ok {
// 		fsData, err := yaml.Marshal(fs)
// 		if err != nil {
// 			return err
// 		}
// 		err = yaml.Unmarshal(fsData, &c.Backend.FileSystem)
// 		if err != nil {
// 			return err
// 		}
// 	} else {
// 		return errors.New("unknown backend type")
// 	}

// 	return nil
// }

func GenerateConfig(path string) (*Config, error) {
	yamlData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config Config
	err = yaml.Unmarshal(yamlData, &config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling YAML: %v", err)
	}
	return &config, nil
}
