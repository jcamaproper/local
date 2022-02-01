package main

import (
	"fmt"
	"log"
	"time"
	"utils/chrome"

	ch "github.com/chromedp/chromedp"
)

const link = "#bt_facebook3"
const b = "#mount_0_0_gF > div > div:nth-child(1) > div > div.rq0escxv.l9j0dhe7.du4w35lb > div > div > div.j83agx80.cbu4d94t.d6urw2fd.dp1hu0rb.l9j0dhe7.du4w35lb > div.l9j0dhe7.dp1hu0rb.cbu4d94t.j83agx80 > div:nth-child(1) > div.rq0escxv.l9j0dhe7.du4w35lb.j83agx80.pfnyh3mw.taijpn5t.gs1a9yip.owycx6da.btwxx1t3.ihqw7lf3.cddn0xzi > div > div > div > div.rq0escxv.l9j0dhe7.du4w35lb.j83agx80.cbu4d94t.pfnyh3mw.d2edcug0.hpfvmrgz.nqmvxvec.ph5uu5jm.b3onmgus.e5nlhep0.ecm0bbzt > div > div > a > div > svg > g > image"

func main() {

	// Create a context
	ctx, cancel, err := chrome.StartChrome(false)
	if err != nil {
		println(err.Error())
		return
	}
	defer cancel()
	fmt.Println(ctx)

	//Oper URL

	url := "https://www.record.com.mx/estadisticas"
	if err := ch.Run(
		ctx,
		ch.Navigate(url)); err != nil {
		println(err.Error())
		return
	}

	if err := ch.Run(
		ctx,
		ch.Click(link, ch.ByQuery)); err != nil {
		println(err.Error())
		return
	}
	f := ""
	ch.Location(&f)

	time.Sleep(5 * time.Second)

	// get the list of the targets
	infos, err := ch.Targets(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if len(infos) == 0 {
		log.Println("no targets")
	}

	// create context attached to the specified target ID.
	// this example just uses the first target,
	// you can search for the one you want.
	tabCtx, cancel := ch.NewContext(ctx, ch.WithTargetID(infos[0].TargetID))
	defer cancel()

	if err := ch.Run(
		tabCtx,
		ch.Click(b, ch.ByQuery)); err != nil {
		fmt.Println("Error", err)

	}

	/* if err := ch.Run(tabCtx, ch.Navigate("https://www.google.com/")); err != nil {
		log.Fatal(err)
	} */

	//time.Sleep(1 * time.Second)
}
