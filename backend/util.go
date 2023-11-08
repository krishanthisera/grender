package backend

import (
	"net/url"
	"path"
	"path/filepath"
	"strings"
)

func generateRelativePath(u string) (string, error) {
	url, err := url.Parse(u)
	if err != nil {
		return "", err
	}
	path := path.Join(url.Host, url.Path)

	return strings.TrimSuffix(path, filepath.Ext(path)), nil
}
