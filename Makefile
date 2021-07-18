GIT_VERSION?=$(shell git describe --tag)

.PHONY: build
build: 
	go build -o bin/gomaze github.com/TerryHowe/gomaze/cmd/gomaze
	pkill gomaze || true
	./bin/gomaze &
	open http://127.0.0.1:3000