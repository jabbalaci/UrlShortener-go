cat:
	cat Makefile

run:
	go run ./cmd/urlshortener

test:
	go test ./... -v

install:
	go install ./cmd/urlshortener
