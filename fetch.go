package main

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"strings"
	"time"
)

const (
	initUrl = "https://www.keishicho-gto.metro.tokyo.lg.jp/keishicho-u/reserve/offerList_detail.action?tempSeq=363"
	target  = "#008A2B" // the green color
)

func fetch() bool {
	rootCtx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// find Chrome binary
	chromePath := os.Getenv("CHROME_PATH")
	if chromePath == "" {
		chromePath = "/usr/bin/chromium-browser"
	}

	// create allocator with necessary flags
	allocCtx, allocCancel := chromedp.NewExecAllocator(rootCtx,
		chromedp.ExecPath(chromePath),
		chromedp.Headless,
		chromedp.DisableGPU,
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-dev-shm-usage", true),
	)
	defer allocCancel()

	// create chrome instance
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// navigate to a page, wait for an element, click
	var (
		found bool
	)

	// run 4 times
	for i := 0; i < 5; i++ {
		var output string

		// the first time is different
		if i == 0 {
			if err := chromedp.Run(ctx,
				chromedp.Navigate(initUrl),
				chromedp.WaitVisible(`body table input[value="2週後＞"]`),
				chromedp.OuterHTML(`body`, &output, chromedp.ByQuery),
			); err != nil {
				log.Printf("Error at the %d fetch: %s\n", i, err)
				return false
			}
		} else {
			if err := chromedp.Run(ctx,
				chromedp.WaitVisible(`body table input[value="2週後＞"]`),
				chromedp.Click(`input[value="2週後＞"]`, chromedp.NodeVisible),
				chromedp.Sleep(2*time.Second),
				chromedp.OuterHTML(`body`, &output, chromedp.ByQuery),
			); err != nil {
				log.Printf("Error at the %d fetch: %s\n", i, err)
				return false
			}
		}

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(output))
		if err != nil {
			log.Fatal(err)
		}
		doc.Find("#height_auto_29の国･地域以外の方で、住民票のない方 td").Each(func(i int, s *goquery.Selection) {
			// Find the text of the element
			v := s.Text()
			// Print the text
			if strings.Contains(v, target) {
				found = true
			}
		})

		if found {
			return true
		}
	}

	return false
}
