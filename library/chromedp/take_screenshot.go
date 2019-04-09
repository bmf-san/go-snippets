package main

import (
	"context"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chromedp/chromedp"
)

type Draw struct {
	Text string
}

func handler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("index.tpl"))

	tpl.Execute(w, Draw{
		Text: "Hello World",
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	srv := &http.Server{
		Addr:    ":9999",
		Handler: mux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Print(err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM)
	go takeScreenShot(ch)
	<-ch

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if err := srv.Shutdown(ctx); err != nil {
		log.Print(err)
	}
}

func takeScreenShot(ch chan os.Signal) {
	defer func() {
		ch <- syscall.SIGTERM
	}()
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	if err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(`http://localhost:9999`),
		chromedp.Sleep(2 * time.Second),
		chromedp.WaitVisible(`#target`, chromedp.ByID),
		chromedp.Screenshot(`#target`, &buf, chromedp.NodeVisible, chromedp.ByID),
	}); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("result.png", buf, 0644); err != nil {
		log.Fatal(err)
	}
}
