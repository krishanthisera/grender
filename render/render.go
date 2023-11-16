package render

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/google/uuid"
)

func (config Config) Render(webAddr string) (*string, error) {
	// Validate the URL
	webAddr = func(val string) string {
		u, err := url.Parse(strings.TrimLeft(val, "/"))
		if err != nil {
			panic(err)
		}
		if u.Scheme == "" {
			u.Scheme = "https"
		}
		return u.String()
	}(webAddr)

	// Create a new Chrome headless instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	var html string
	// Set request headers
	headers := func() network.Headers {
		h := network.Headers{"X-Grender-Request-Id": fmt.Sprintf("%v", uuid.New())}
		for _, v := range config.RequestHeaders {
			h[v.Name] = v.Value
		}
		fmt.Println(h)
		return h
	}()

	if err := chromedp.Run(ctx, pageRender(webAddr, config.PageWailCondition, time.Duration(config.PageWaitTime*float32(time.Second)), &headers, &html)); err != nil {
		return nil, err
	}
	return &html, nil
}

// This is the function that does the actual rendering
func pageRender(webAddr string, waitCondition string, pageWaitTime time.Duration, headers *network.Headers, html *string) chromedp.Tasks {
	return chromedp.Tasks{
		network.Enable(),
		network.SetExtraHTTPHeaders(*headers),
		chromedp.Navigate(webAddr),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var result bool
			startTime := time.Now()
			for time.Since(startTime) < pageWaitTime {
				if err := chromedp.Evaluate(waitCondition, &result).Do(ctx); err != nil {
					return err
				}
				if result {
					break // The condition is met, exit the loop.
				}

				// Sleep for a short duration before re-evaluating the condition.
				time.Sleep(500 * time.Millisecond)
			}

			if !result {
				return fmt.Errorf("timeout [%v] waiting for window.prerenderReady to become true", pageWaitTime)
			}

			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			*html, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			return err
		}),
	}
}
