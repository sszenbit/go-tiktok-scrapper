package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

var wg sync.WaitGroup

func main() {
	const reqCount = 10
	wg.Add(reqCount)
	for i := 0; i < reqCount; i++ {
		go func() {
			defer wg.Done()
			scrapeVideoCount()
		}()
	}
	wg.Wait()
}

var failed = 0

func scrapeVideoCount() {
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting", r.URL)
		r.Headers.Set("User-Agent", RandomString())
	})

	c.OnError(func(_ *colly.Response, err error) {
		failed++
	})

	c.OnHTML(`[data-e2e="music-video-count"]`, func(e *colly.HTMLElement) {
		fmt.Println(strings.Split(e.Text, " ")[0])
	})

	c.Visit("https://vm.tiktok.com/ZMLEe7HnQ/")
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
