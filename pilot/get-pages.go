package pilot

import "github.com/krishanthisera/grender/backend"

func GetPages(url string) ([]byte, error) {
	// GetPages logic
	fs := backend.FileSystem{BaseDir: "/tmp/"}

	res, err := backend.Backend.Get(fs, url)
	if err != nil {
		return nil, err
	}
	return res, nil

}
