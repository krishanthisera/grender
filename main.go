package main

import (
	"fmt"

	"github.com/krishanthisera/grender/backend"
	"github.com/krishanthisera/grender/render"
)

func main() {
	url := "https://www.yd.com.au/au/white-striker-sneaker-y233fc11"

	pageWaitCondition := `(function() {
		return window.prerenderReady === true;
})()
`

	result, err := render.RenderingConfigs{PageWailCondition: pageWaitCondition, PageWaitTime: 16}.Render(url)

	if err != nil {
		panic(err)
	}
	// log.Println(string(result))

	// s3 := backend.S3{BucketName: "s3 test"}
	fs := backend.FileSystem{BaseDir: "/tmp/"}
	// backend.Backend.Put(s3, url, result)

	fmt.Println(result)

	if err := backend.Backend.Put(fs, url, []byte(*result)); err != nil {
		panic(err)
	}
}
