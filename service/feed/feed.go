package feed

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/NoHomey/chaos-go-camp-proj/ctxerr"
	"github.com/mmcdole/gofeed"
	"go.uber.org/zap"
)

//ErrInvalURL is used to report invalid URLs.
type ErrInvalURL struct {
	wrapped error
	URL     string
}

func (err ErrInvalURL) Error() string {
	return fmt.Sprintf("Invalid URL: %s. Error: %s", err.URL, err.wrapped.Error())
}

func (err ErrInvalURL) Unwrap() error {
	return err.wrapped
}

//Text returns human readable error text.
func (err ErrInvalURL) Text() string {
	return fmt.Sprintf("Invalid URL: %s", err.URL)
}

//Context returns error Context.
func (err ErrInvalURL) Context() ctxerr.Context {
	return ctxerr.Context{
		Name: "invalid-url",
		Data: err.URL,
	}
}

//HttpStatusCode returns http status code for the error.
func (err ErrInvalURL) HttpStatusCode() int {
	return 400
}

//ErrParseFeed is used to report errors from parsing feeds.
type ErrParseFeed struct {
	wrapped error
	URL     string
}

func (err ErrParseFeed) Error() string {
	return fmt.Sprintf("Could not parse feed for URL: %s. Error: %s", err.URL, err.wrapped.Error())
}

func (err ErrParseFeed) Unwrap() error {
	return err.wrapped
}

//Text returns human readable error text.
func (err ErrParseFeed) Text() string {
	return fmt.Sprintf("Could not parse feed for URL: %s", err.URL)
}

//Context returns error Context.
func (err ErrParseFeed) Context() ctxerr.Context {
	return ctxerr.Context{
		Name: "parse-feed",
		Data: err.URL,
	}
}

//HttpStatusCode returns http status code for the error.
func (err ErrParseFeed) HttpStatusCode() int {
	return 400
}

//Details represents feed details extracted from given URL.
type Details struct {
	URL         string
	Title       string
	Author      string
	Description string
}

//Item represents feed entry/item.
type Item struct {
	URL         string
	Title       string
	Description string
	PublishedAt time.Time
}

//Service abstracts the feed service.
type Service interface {
	Details(ctx context.Context, url string) (*Details, ctxerr.Error)
	FeedBatches(ctx context.Context, url string, batchSize uint) ([][]*Item, ctxerr.Error)
}

//Use constructs Service wich uses the given logger.
func Use(logger *zap.Logger) Service {
	return service{logger}
}

type service struct {
	logger *zap.Logger
}

func (srvc service) Details(ctx context.Context, url string) (*Details, ctxerr.Error) {
	feed, cerr := srvc.fetchFeed(ctx, url)
	if cerr != nil {
		return nil, cerr
	}
	author := ""
	if feed.Author != nil {
		author = feed.Author.Name
	}
	return &Details{
		URL:         url,
		Title:       feed.Title,
		Author:      author,
		Description: feed.Description,
	}, nil
}

func (srvc service) FeedBatches(ctx context.Context, url string, batchSize uint) ([][]*Item, ctxerr.Error) {
	feed, cerr := srvc.fetchFeed(ctx, url)
	if cerr != nil {
		return nil, cerr
	}
	batches := [][]*Item{
		make([]*Item, batchSize),
	}
	batch := uint(0)
	currSize := uint(0)
	for _, item := range feed.Items {
		if currSize == batchSize {
			batches = append(batches, make([]*Item, batchSize))
			batch++
			currSize = 0
		}
		publishedAt := time.Now()
		if item.PublishedParsed != nil {
			publishedAt = *item.PublishedParsed
		}
		batches[batch][currSize] = &Item{
			URL:         item.Link,
			Title:       item.Title,
			Description: item.Description,
			PublishedAt: publishedAt,
		}
		currSize++
	}
	return batches, nil
}

func (srvc service) fetchFeed(ctx context.Context, urlStr string) (*gofeed.Feed, ctxerr.Error) {
	url, err := url.Parse(urlStr)
	if err != nil {
		srvc.logger.Error(
			"Failed to parse url",
			zap.String("url", urlStr),
			zap.Error(err),
		)
		return nil, ErrInvalURL{err, urlStr}
	}
	req, err := http.NewRequestWithContext(ctx, "GET", url.String(), nil)
	if err != nil {
		srvc.logger.Error(
			"Failed to create new request",
			zap.String("url", url.String()),
			zap.Error(err),
		)
		return nil, ctxerr.NewInternal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		srvc.logger.Error(
			"Failed to fetch feed",
			zap.String("url", url.String()),
			zap.Error(err),
		)
		return nil, ctxerr.NewInternal(err)
	}
	defer resp.Body.Close()
	feed, err := gofeed.NewParser().Parse(resp.Body)
	if err != nil {
		srvc.logger.Error(
			"Failed to parse feed",
			zap.String("url", url.String()),
			zap.Error(err),
		)
		return nil, ErrParseFeed{err, url.String()}
	}
	srvc.logger.Info(
		"Fetched and parsed feed",
		zap.String("url", url.String()),
	)
	return feed, nil
}
