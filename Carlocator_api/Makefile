.PHONY: build clean deploy gomodgen

build:
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/carlocator cmd/main.go

clean:
	rm -rf ./bin ./vendor go.sum

deploy: clean build
	sls deploy --verbose