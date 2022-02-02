package main

import (
	"fmt"
	"log"
	"time"

	//"time"
	"utils/chrome"

	ch "github.com/chromedp/chromedp"
)

const (
	instaLogginButton = "#loginForm > a"
	recordInstaButton = "#bt_instagram3"
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

	//click to open new tab
	if err := ch.Run(
		ctx,
		ch.Click(recordInstaButton, ch.ByQuery),
	); err != nil {
		println(err.Error())
		return
	}

	defer cancel()

	//time.Sleep(20 * time.Second)
	// get the list of the targets first context
	infos, err := ch.Targets(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if len(infos) == 0 {
		log.Println("no targets")
	}

	//generate new context from first targetID context1
	tabCtx, cancel := ch.NewContext(ctx, ch.WithTargetID(infos[0].TargetID))
	defer cancel()

	//get URL Second tab
	if err := ch.Run(
		tabCtx,
		// Get Current URL
		ch.Location(&vurl2)); err != nil {
		println(err.Error())
		return
	}

	fmt.Println("The second URL is: ", vurl2)

	if err := ch.Run(
		tabCtx,
		// Get Current URL
		ch.Click(instaLogginButton, ch.ByQuery),
	); err != nil {
		println(err.Error())
		return
	}

	time.Sleep(20 * time.Second)
}
