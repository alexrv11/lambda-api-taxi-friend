.PHONY: deps clean build rebuild

deps:
	go get -u ./...

clean: 
	rm -rf ./dist
	
build:
	GOOS=linux GOARCH=amd64 go build -o dist/ping ./lambdas/ping
	GOOS=linux GOARCH=amd64 go build -o dist/drivers ./lambdas/drivers

rebuild:clean build
