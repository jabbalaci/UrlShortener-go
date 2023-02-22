package main

import (
	"fmt"
	"log"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	url := "https://index.hu"

	pngFile := "/tmp/qr.png"
	err := qrcode.WriteFile(url, qrcode.Medium, 512, pngFile)
	if err != nil {
		log.Printf("createQrCode: %v\n", err)
	}
	fmt.Println("END")
}
