package pilot

import (
	"github.com/krishanthisera/grender/backend"
)

func (c *renderAndCacheConfig) RenderAndCache(url string) ([]byte, error) {

	res, err := backend.Backend.Get(*c.backend, url)

	// If errored te app must render the page on the fly
	if err != nil {
		page, err := c.render.Render(url)
		if err != nil {
			return nil, err
		}
		// If the page is rendered successfully, save it to the backend
		if err := backend.Backend.Put(*c.backend, url, []byte(*page)); err != nil {
			return []byte(*page), err
		}
		return []byte(*page), nil
	}

	return res, nil
}
