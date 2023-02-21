package bitly

import (
	"os"
	"strings"

	"github.com/retgits/bitly/client"
	"github.com/retgits/bitly/client/bitlinks"
)

func Shorten(url string) (string, error) {
	bitly := client.NewClient().WithAccessToken(os.Getenv("BITLY_ACCESS_TOKEN"))
	service := bitlinks.New(bitly)
	short, err := service.ShortenLink(&bitlinks.ShortenRequest{LongURL: url})
	if err != nil {
		return "", err
	}
	return short.Link, nil
}

func Expand(url string) (string, error) {
	bitly := client.NewClient().WithAccessToken(os.Getenv("BITLY_ACCESS_TOKEN"))
	url = strings.TrimPrefix(url, "https://")
	service := bitlinks.New(bitly)
	result, err := service.ExpandBitlink(bitlinks.Link{BitlinkID: url})
	if err != nil {
		return "", err
	}
	return result.LongURL, nil
}
