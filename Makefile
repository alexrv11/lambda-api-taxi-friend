.PHONY: deps clean build rebuild

deps:
	go get -u ./...

clean: 
	rm -rf ./dist
	
build:
	GOOS=linux GOARCH=amd64 go build -o dist/ping/ping ./lambdas/ping

rebuild:clean build
