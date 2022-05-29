package scrapperpc

import (
	context "context"
	"log"
	"math/rand"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type ScrapperServer struct {
}

func (s *ScrapperServer) GetTrackInfo(ctx context.Context, req *TrackInfoRequest) (*TrackInfoResponse, error) {
	url := req.GetUrl()
	videoCount, err := scrapeVideoCount(url)
	if err != nil {
		return nil, err
	}
	return &TrackInfoResponse{Info: &TrackInfo{
		TrackTiktokId:     "1",
		TrackTitle:        "2",
		TrackPlayUrl:      "3",
		TrackCoverUrl:     "4",
		TrackAuthorName:   "5",
		InitialVideoCount: videoCount,
	}}, nil
}

func scrapeVideoCount(url string) (int64, error) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", RandomString())
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	resCh := make(chan int64)
	errCh := make(chan error)
	c.OnHTML(`[data-e2e="music-video-count"]`, func(e *colly.HTMLElement) {
		parsedValue, err := strconv.ParseInt(strings.Split(e.Text, " ")[0], 10, 64)
		if err != nil {
			errCh <- err
			return
		}
		resCh <- parsedValue

	})
	c.Visit(url)
	select {
	case err := <-errCh:
		return 0, err
	case res := <-resCh:
		return res, nil
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
