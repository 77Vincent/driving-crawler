package main

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"log"
	"strings"
	"time"
)

const (
	initUrl = "https://www.keishicho-gto.metro.tokyo.lg.jp/keishicho-u/reserve/offerList_detail.action?tempSeq=363"
	target  = "#008A2B" // the green color
)

func fetch() bool {
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 20*time.Second)
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
				log.Fatal(err)
			}
		} else {
			if err := chromedp.Run(ctx,
				chromedp.WaitVisible(`body table input[value="2週後＞"]`),
				chromedp.Click(`input[value="2週後＞"]`, chromedp.NodeVisible),
				chromedp.Sleep(2*time.Second),
				chromedp.OuterHTML(`body`, &output, chromedp.ByQuery),
			); err != nil {
				log.Fatal(err)
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
