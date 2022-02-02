package main

import (
	"fmt"
	"time"
	"utils/chrome"

	ch "github.com/chromedp/chromedp"
)

const (
	luckyButton  = "body > div.L3eUgb > div.o3j99.ikrT4e.om7nvf > form > div:nth-child(1) > div.A8SBwf > div.FPdoLc.lJ9FBc > center > input.RNmpXc"
	mobileButton = "#agWOA9yL7u > div > div.masthead-site-nav-container.js-nano-container > nav > ul.masthead-nav-topics > li:nth-child(3) > a"
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

	url := "https://www.google.com.mx/"
	//url2 := "https://www.xataka.com/"

	var vurl1 string
	var vurl2 string

	if err := ch.Run(
		ctx,
		ch.Navigate(url),
		// Get Current URL
		ch.Location(&vurl1)); err != nil {
		println(err.Error())
		return
	}

	fmt.Println("The current URL is: ", vurl1)

	//open second tab -- same broser
	ctx2, _ := ch.NewContext(ctx)

	// ensure the second tab is created
	if err := ch.Run(ctx2); err != nil {
		panic(err)
	}

	if err := ch.Run(
		ctx2,
		ch.Navigate(url),
		// Get Current URL
		ch.Location(&vurl2)); err != nil {
		println(err.Error())
		return
	}

	fmt.Println("The current URL is: ", vurl2)
	time.Sleep(10 * time.Second)
	if err := ch.Run(
		ctx2,
		ch.Click(luckyButton, ch.ByQuery),
	); err != nil {
		println(err.Error())
		return
	}
	time.Sleep(10 * time.Second)
	if err := ch.Cancel(ctx2); err != nil {
		println(err.Error())
		return
	}

	time.Sleep(20 * time.Second)
	// get the list of the targets first context
	/* infos, err := ch.Targets(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if len(infos) == 0 {
		log.Println("no targets")
	}

	//click to open second tab on the same browser
	if err := ch.Run(
		ctx,
		ch.Click(link, ch.ByQuery)); err != nil {
		println(err.Error())
		return
	}

	// get the list of the targets
	infos2, err := ch.Targets(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if len(infos) == 0 {
		log.Println("no targets")
	}

	fmt.Println(infos, infos2)

	// create context attached to the specified target ID.
	// this example just uses the first target,
	// you can search for the one you want.
	tabCtx, cancel := ch.NewContext(ctx, ch.WithTargetID(infos2[0].TargetID))
	defer cancel()

	if err := ch.Run(
		tabCtx,
		ch.Click(b, ch.ByQuery)); err != nil {
		fmt.Println("Error", err)

	} */

	/* if err := ch.Run(tabCtx, ch.Navigate("https://www.google.com/")); err != nil {
		log.Fatal(err)
	} */

	//time.Sleep(1 * time.Second)
}
