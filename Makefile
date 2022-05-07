.PHONY: build clean deploy

build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/get-skills get-skills/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/save-skills save-skills/main.go

	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/get-profile get-profile/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/save-profile save-profile/main.go

	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/get-presigned-url get-presigned-url/main.go

	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/authorizer authorizer/main.go

	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/get-languages get-languages/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/save-languages save-languages/main.go


	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/get-educations get-educations/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/save-educations save-educations/main.go

	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/get-projects get-projects/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/save-projects save-projects/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
