package main

import (
	"fmt"
	"utils/chrome"

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

	propertyLabel = "PropertyCode_LookupCode"
	chargeCode    = "#ChargeCode_LookupCode"
	habMonth      = "#HAPMonth_TextBox"
	submitButton  = "#Submit_Button"

	password = "rwt3fxq@egb3YEG*zvn"
	username = "martha"
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

	//Enter consolidated report data
	//***** Can not perform actions inside iframe*******
	//var ok bool
	if err := ch.Run(
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
		ch.Click(submitButton, ch.ByJSPath),
	); err != nil {
		fmt.Println("Error", err)
	}
	//ch.EvaluateAsDevTools(fmt.Sprintf(`document.querySelector("%s").dispatchEvent(new Event("change"))`, securityQuestionInput), &ok),
	//go to add receipts
	fmt.Println("The current URL is: ", url1)

}
