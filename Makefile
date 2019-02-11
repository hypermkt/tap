ASSETS_DIR=assets/
BUILD_FILE=bin/tap

assets: statik/statik.go
	go get github.com/rakyll/statik
	statik -src=$(ASSETS_DIR)

build: bin/tap
	go build -o $(BUILD_FILE)
