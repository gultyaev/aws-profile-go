.PHONY: build clean deploy

build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/get-skills get-skills/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/save-skills save-skills/main.go
#	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/get-profile get-profile/main.go
#	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/save-profile save-profile/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
