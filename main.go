package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/cdp"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

func main() {
	var url string
	fmt.Print("Enter the link: ")
	fmt.Fscan(os.Stdin, &url)

	browser := rod.New().Client(CreateClient())
	if err := browser.Connect(); err != nil {
		log.Fatalln("error connection:", err)
	}

	page, err := browser.Page(proto.TargetCreateTarget{URL: url}) // URL
	if err != nil {
		log.Fatalln("error create page:", err)
	}

	page.MustWaitStable() // Loading...
	time.Sleep(10 * time.Second)

	err = page.Mouse.MoveTo(proto.NewPoint(450+(rand.Float64()*10), 380))
	if err != nil {
		log.Fatalln("error mouse move:", err)
	}

	go func() {
		fmt.Println("Launch of mining...")
		for {
			page.Mouse.MustClick(proto.InputMouseButtonLeft) // Start click
			page.MustScreenshot("status.png")
			fmt.Println("You have started mining!")
			time.Sleep(25 * time.Minute)

			page.Mouse.MustClick(proto.InputMouseButtonLeft) // Stop click
			page.MustScreenshot("status.png")
			fmt.Println("You have stopped mining!")
			time.Sleep(90 * time.Minute)
		}
	}()

	bufio.NewReader(os.Stdin).ReadBytes('\n')
	browser.Close()
}

func CreateClient() *cdp.Client {
	ctx := context.Background()
	launcher := launcher.New()
	launcher.NoSandbox(true)
	u, err := launcher.Context(ctx).Launch()
	if err != nil {
		log.Fatalln("error launch loader:", err)
	}

	client, err := cdp.StartWithURL(ctx, u, nil)
	if err != nil {
		log.Fatalln("error starting launcher")
	}

	return client
}
