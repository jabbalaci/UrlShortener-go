package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {
	urlString := "https://www.amazon.com/Concurrency-Go-Tools-Techniques-Developers/dp/1491941197/ref=sr_1_1?crid=1CFL729O80W6B&keywords=concurrency+in+go&qid=1677268474&s=books&sprefix=Concurrency+in+Go%2Cstripbooks-intl-ship%2C170&sr=1-1"

	decodedURLString, err := url.QueryUnescape(urlString)
	if err != nil {
		panic(err)
	}
	decodedURLString = strings.ReplaceAll(decodedURLString, " ", "+")

	fmt.Println(decodedURLString)
}
