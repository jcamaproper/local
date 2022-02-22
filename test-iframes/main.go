package main

import (
	"fmt"
	"log"
	"time"
	"utils/chrome"

	"github.com/chromedp/cdproto/cdp"
	ch "github.com/chromedp/chromedp"
)

const (
	userLabel   = "#Username"
	passLabel   = "#Password_Text"
	loginButton = "#cmdLogin1"
	environment = "#Destination"

	menuSelector        = "#miSearch > input[type=text]"
	basicDataMenu       = "#miS0-4 > a"
	consolidatedReports = "#miS3-8-1 > a"

	//Consolidated Receipt Page Elements
	consolidateIframe = "#filter"
	propertyLabel     = "PropertyCode_LookupCode"
	chargeCode        = "#ChargeCode_LookupCode"
	habMonth          = "#HAPMonth_TextBox"
	submitButton      = "#Submit_Button"

	//Affordable Consolidated Receip Page Elements
	affordableIframe = "#filter"
	checkNumberField = "#CheckNumber_TextBox"
	postDateField    = "#PostDate_TextBox"
	postMonthField   = "#PostMonth_TextBox"
	saveButton       = "#Save_Button"

	//Credentials
	password = ""
	username = ""
)

func main() {

	// Create a context
	ctx, cancel, err := chrome.StartChrome(false)
	if err != nil {
		println(err.Error())
		return
	}
	defer cancel()

	//Open URL

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
		ch.SetValue(environment, "Test", ch.ByQuery),
		ch.Click(loginButton, ch.ByQuery),
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

	//*************GET IFRAME IN NODES*************
	//Enter consolidated report data

	//var iframes, forms []*cdp.Node
	time.Sleep(2 * time.Second)
	var iframes []*cdp.Node
	if err := ch.Run(
		ctx,
		ch.WaitReady(consolidateIframe, ch.ByID),
		ch.Nodes(consolidateIframe, &iframes, ch.ByQuery),
	); err != nil {
		log.Fatal(err)
	}
	if err := ch.Run(
		ctx,
		ch.WaitReady(propertyLabel, ch.ByID, ch.FromNode(iframes[0])),
		ch.Click(propertyLabel, ch.ByID, ch.FromNode(iframes[0])),
		ch.KeyEvent("all-rpm"),
		ch.Click(chargeCode, ch.ByID, ch.FromNode(iframes[0])),
		ch.KeyEvent("subsidy"),
		ch.Click(habMonth, ch.ByID, ch.FromNode(iframes[0])),
		ch.KeyEvent("01/2022"),
		ch.Click(submitButton, ch.ByID, ch.FromNode(iframes[0])),
	); err != nil {
		log.Fatal(err)
	}

	//*************GET IFRAME IN NODES*************

	//*************GET IFRAME IN NODES*************
	//Affordable Consolidated Receipt
	time.Sleep(2 * time.Second)
	if err := ch.Run(
		ctx,
		ch.WaitReady(affordableIframe, ch.ByID),
		ch.Nodes(affordableIframe, &iframes, ch.ByQuery),
	); err != nil {
		log.Fatal(err)
	}
	if err := ch.Run(
		ctx,
		ch.WaitReady(checkNumberField, ch.ByID, ch.FromNode(iframes[0])),
		ch.Click(checkNumberField, ch.ByID, ch.FromNode(iframes[0])),
		ch.KeyEvent("123456"),
		ch.Click(postDateField, ch.ByID, ch.FromNode(iframes[0])),
		ch.KeyEvent("02/03/2022"),
		ch.Click(postMonthField, ch.ByID, ch.FromNode(iframes[0])),
		ch.KeyEvent("02"),
		//ch.Click(saveButton, ch.ByID, ch.FromNode(iframes[0])),
	); err != nil {
		log.Fatal(err)
	}

	//******************DO NOT WORK, The infos detected only one target, no an additional for the iFrame****************
	/* // get the list of the targets first context
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
	fmt.Println(tabCtx) */

	//Enter consolidated report data
	//***** Can not perform actions inside iframe*******
	//var ok bool
	/* 	if err := ch.Run(
	   		ctx,
	   		ch.WaitReady(propertyLabel, ch.ByJSPath),
	   		//ch.SetValue(propertyLabel, "blvd", ch.ByQuery),
	   		//ch.SetValue(chargeCode, "subsidy", ch.ByQuery),
	   		//ch.SetValue(habMonth, "01/2022", ch.ByQuery),
	   		//ch.Click(propertyLabel, ch.ByID),
	   		ch.SetAttributeValue(propertyLabel, "value", "texto prueba", ch.ByJSPath),
	   		//ch.Blur(propertyLabel, ch.ByQuery),
	   		//ch.KeyEvent("blvd"),
	   		//ch.EvaluateAsDevTools(fmt.Sprintf(`document.querySelector("%s").value="texto prueba"`, propertyLabel), &ok),
	   		//ch.Click(chargeCode, ch.ByQuery),
	   		//ch.KeyEvent("subsidy"),
	   		//ch.Click(habMonth, ch.ByQuery),
	   		//ch.KeyEvent("01/2022"),
	   		//ch.Click(submitButton, ch.ByJSPath),
	   	); err != nil {
	   		fmt.Println("Error", err)
	   	}
	   	//ch.EvaluateAsDevTools(fmt.Sprintf(`document.querySelector("%s").dispatchEvent(new Event("change"))`, securityQuestionInput), &ok),
	   	//go to add receipts
	   	fmt.Println("The current URL is: ", url1) */

}
