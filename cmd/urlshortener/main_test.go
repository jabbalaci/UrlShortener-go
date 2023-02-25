package main

import (
	"testing"

	"github.com/jabbalaci/UrlShortener-go/lib/bitly"
	"github.com/stretchr/testify/assert"
)

/////////////////////////////////////////////////////////////////////////////

func Test1(t *testing.T) {
	original_url := "https://www.reddit.com"
	short_url, _ := bitly.Shorten(original_url)
	assert.Equal(t, short_url, "https://bit.ly/3KmRGIy")
	expanded_url, _ := bitly.Expand(short_url)
	assert.Equal(t, expanded_url, "https://www.reddit.com/")
	assert.Equal(t, fuzzy_match(original_url, expanded_url), true)

}

func Test2(t *testing.T) {
	original_url := "https://www.reddit.com/"
	short_url, _ := bitly.Shorten(original_url)
	assert.Equal(t, short_url, "https://bit.ly/3KmRGIy")
	expanded_url, _ := bitly.Expand(short_url)
	assert.Equal(t, expanded_url, "https://www.reddit.com/")
	assert.Equal(t, fuzzy_match(original_url, expanded_url), true)
}

func Test3(t *testing.T) {
	url1 := "https://www.reddit.com"
	url2 := "https://www.reddit.com"
	assert.Equal(t, fuzzy_match(url1, url2), true)
	//
	url1 = "https://www.reddit.com/"
	url2 = "https://www.reddit.com/"
	assert.Equal(t, fuzzy_match(url1, url2), true)
	//
	url1 = "https://www.reddit.com/"
	url2 = "https://www.reddit.com"
	assert.Equal(t, fuzzy_match(url1, url2), true)
	//
	url1 = "https://www.reddit.com"
	url2 = "https://www.reddit.com/"
	assert.Equal(t, fuzzy_match(url1, url2), true)
}

func Test4(t *testing.T) {
	url1 := "https://www.amazon.com/Concurrency-Go-Tools-Techniques-Developers/dp/1491941197/ref=sr_1_1?crid=FQ1NOP9U8RX0&keywords=concurrent+go&qid=1677335205&sprefix=%2Caps%2C154&sr=8-1"
	url2 := "https://www.amazon.com/Concurrency-Go-Tools-Techniques-Developers/dp/1491941197/ref=sr_1_1?crid=FQ1NOP9U8RX0&keywords=concurrent+go&qid=1677335205&sprefix=,aps,154&sr=8-1"
	assert.Equal(t, fuzzy_match(url1, url2), true)
}
