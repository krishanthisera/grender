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
	headers := map[string]interface{}{
		"Grerender-Request-Id": "my request header",
	}

	if err := chromedp.Run(ctx, pageRender(webAddr, config.PageWailCondition, time.Duration(config.PageWaitTime*float32(time.Second)), headers, &html)); err != nil {
		return nil, err
	}
	return &html, nil
}

// This is the function that does the actual rendering
func pageRender(webAddr string, pageWaitCondition string, pageWaitTime time.Duration, headers network.Headers, html *string) chromedp.Tasks {
	return chromedp.Tasks{
		network.Enable(),
		// Setting the netowrk headers has been disable as: ..."net::ERR_FAILED","canceled":false,"corsErrorStatus":{"corsError":"HeaderDisallowedByPreflightResponse","failedParameter":"grerender-request-id"}}
		// network.SetExtraHTTPHeaders(network.Headers(headers)),
		chromedp.Navigate(webAddr),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var result bool
			startTime := time.Now()
			for time.Since(startTime) < pageWaitTime {
				if err := chromedp.Evaluate(pageWaitCondition, &result).Do(ctx); err != nil {
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
