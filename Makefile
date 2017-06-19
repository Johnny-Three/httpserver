PKG_NAME=$(shell basename `pwd`)

hi:
	$(PKG_NAME)

run: 
	cd ./bin && ./httpserver

build: 
	go build -o ./bin/$(PKG_NAME) -v ./src/ 

test:
	go test -v ./src/...

bench:
	go test ./src/... -bench . -benchmem
