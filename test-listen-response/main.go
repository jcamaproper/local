package main

import (
	"context"
	"fmt"
	"regexp"
	"utils/chrome"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	ch "github.com/chromedp/chromedp"
)

func main() {
	// Create a context
	ctx, cancel, err := chrome.StartChrome(false)
	if err != nil {
		println(err.Error())
		return
	}
	defer cancel()

	//Oper URL

	url := "https://www.record.com.mx/"

	var creditsUrlRegex = regexp.MustCompile(`https://www.record.com.mx/`)
	var url1 string

	if err := ch.Run(
		ctx,
		ch.Navigate(url),
		// Get Current URL
		ch.Location(&url1)); err != nil {
		println(err.Error())
		return
	}

	c, c1 := ListenRequests(ctx, creditsUrlRegex)

	fmt.Println(c, c1)

}

func ListenRequests(ctx context.Context, urlRegexp *regexp.Regexp) (chan []byte, chan error) {
	result := make(chan []byte)
	fail := make(chan error)

	ch.ListenTarget(ctx, func(ev interface{}) {
		evt, ok := ev.(*network.EventResponseReceived)
		if !ok || (evt.Type != network.ResourceTypeXHR && evt.Type != network.ResourceTypeFetch) {
			return
		}

		if !urlRegexp.MatchString(evt.Response.URL) {
			return
		}

		go func() {
			c := ch.FromContext(ctx)
			body, err := network.GetResponseBody(evt.RequestID).Do(cdp.WithExecutor(ctx, c.Target))
			if err != nil {
				fail <- err
			}
			result <- body
		}()
	})

	return result, fail

}
