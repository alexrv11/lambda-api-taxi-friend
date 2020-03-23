.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./lambdas/ping/ping
	
build:
	GOOS=linux GOARCH=amd64 go build -o lambdas/ping/ping ./lambdas/ping