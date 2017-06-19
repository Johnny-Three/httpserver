PKG_NAME=$(shell basename `pwd`)

hi:
	$(PKG_NAME)

run: 
	cd ./bin && ./httpserver

build: 
	go build -o ./bin/$(PKG_NAME) -v ./src/ 

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/$(PKG_NAME)_lux  -v ./src/

test:
	go test -v ./src/...

bench:
	go test ./src/... -bench . -benchmem
