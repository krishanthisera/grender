package render

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
)

func (config RenderingConfigs) Render(url string) (*string, error) {
	// Create a new Chrome headless instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	var html string
	if err := chromedp.Run(ctx, pageRender(url, config.PageWailCondition, time.Duration(config.PageWaitTime*float32(time.Second)), &html)); err != nil {
		return nil, err
	}
	return &html, nil
}

func pageRender(url string, waitCondition string, pageWaitTime time.Duration, html *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// Wait Condition

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
				return fmt.Errorf("timeout waiting for window.prerenderReady to become true")
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
