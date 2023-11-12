package pilot

import (
	"github.com/krishanthisera/grender/backend"
	"github.com/krishanthisera/grender/render"
)

func GetPages(url string) ([]byte, error) {
	// GetPages logic
	fs := backend.FileSystem{BaseDir: "/tmp/"}

	res, err := backend.Backend.Get(fs, url)

	pageWaitCondition := `(function() {
		return window.prerenderReady === true;
	})()`

	// If errored te app must render the page on the fly
	if err != nil {
		page, err := render.RenderingConfigs{PageWailCondition: pageWaitCondition, PageWaitTime: 16}.Render(url)
		if err != nil {
			return nil, err
		}
		// If the page is rendered successfully, save it to the backend
		if err := backend.Backend.Put(fs, url, []byte(*page)); err != nil {
			return []byte(*page), err
		}
		return []byte(*page), nil
	}

	return res, nil
}
