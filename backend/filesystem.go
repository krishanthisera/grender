package backend

import (
	"os"
	"path/filepath"
)

func (f FileSystem) Put(u string, data []byte) error {
	p, err := generateRelativePath(u)
	if err != nil {
		return err
	}

	// Getting directory path
	dir := filepath.Dir(filepath.Join(f.BaseDir, p))
	os.MkdirAll(dir, os.ModePerm)

	// Writing file to the system
	file, err := os.Create(filepath.Join(f.BaseDir, p) + ".html")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	defer file.Close()
	return nil

}

func (f FileSystem) Get(url string) ([]byte, error) {
	// FileSystem get logic
	p, err := generateRelativePath(url)
	if err != nil {
		return nil, err
	}

	b, err := os.ReadFile(filepath.Join(f.BaseDir, p) + ".html")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (f FileSystem) Delete(url string) error {
	// FileSystem delete logic
	p, err := generateRelativePath(url)
	if err != nil {
		return err
	}

	err = os.Remove(filepath.Join(f.BaseDir, p) + ".html")
	if err != nil {
		return err
	}
	return nil
}
