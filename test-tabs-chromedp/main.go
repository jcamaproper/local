package main

import (
	"fmt"
	"utils/chrome"

	ch "github.com/chromedp/chromedp"
)

func cmain() {

	// Create a context
	ctx, cancel, err := chrome.StartChrome(false)
	if err != nil {
		println(err.Error())
		return
	}
	defer cancel()

	//Oper URL

	url := "https://www.yardiaspnc7.com/90927hrd/pages/LoginAdvanced.aspx"

	var url1 string

	if err := ch.Run(
		ctx,
		ch.Navigate(url),
		// Get Current URL
		ch.Location(&url1)); err != nil {
		println(err.Error())
		return
	}

	if err := ch.Run(
		ctx,
		ch.WaitVisible(userLabel, ch.ByQuery),
		ch.Click(userLabel, ch.ByQuery),
		ch.KeyEvent(username),
		ch.Click(passLabel, ch.ByQuery),
		ch.KeyEvent(password),
		ch.SetValue(enviroment, "Test", ch.ByQuery),
		ch.Click(logginButton, ch.ByQuery),
	); err != nil {
		fmt.Println("Error", err)

	}

	//go to Basic iData Menu
	if err := ch.Run(
		ctx,
		ch.WaitVisible(menuSelector, ch.ByQuery),
		ch.Click(menuSelector, ch.ByQuery),
		ch.KeyEvent("Basic iData Menu"),
		ch.WaitReady(basicDataMenu, ch.ByQuery),
		ch.Click(basicDataMenu, ch.ByQuery),
	); err != nil {
		fmt.Println("Error", err)
	}

	//go to Consolidated
	if err := ch.Run(
		ctx,
		ch.WaitVisible(menuSelector, ch.ByQuery),
		ch.Click(menuSelector, ch.ByQuery),
		ch.KeyEvent("Consolidated"),
		ch.WaitReady(consolidatedReports, ch.ByQuery),
		ch.Click(consolidatedReports, ch.ByQuery),
	); err != nil {
		fmt.Println("Error", err)
	}

	//Enter consolidated report
	if err := ch.Run(
		ctx,
		ch.WaitVisible(menuSelector, ch.ByQuery),
		ch.Click(menuSelector, ch.ByQuery),
		ch.KeyEvent("Consolidated"),
		ch.WaitReady(consolidatedReports, ch.ByQuery),
		ch.Click(consolidatedReports, ch.ByQuery),
	); err != nil {
		fmt.Println("Error", err)
	}

	//go to add receips
	fmt.Println("The current URL is: ", url1)

	/* 	//open second tab -- same broser
	   	ctx2, _ := ch.NewContext(ctx)

	   	// ensure the second tab is created
	   	if err := ch.Run(ctx2); err != nil {
	   		panic(err)
	   	} */

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
