[![Build Status](https://travis-ci.org/alexrv11/lambda-api-taxi-friend.svg?branch=master)](https://travis-ci.org/alexrv11/lambda-api-taxi-friend)
[![codecov](https://codecov.io/gh/alexrv11/lambda-api-taxi-friend/branch/master/graph/badge.svg)](https://codecov.io/gh/alexrv11/lambda-api-taxi-friend)

# Taxifriend server

I known, there is no document in this repo so in the next week I will improve all the weak points.
- CI
- build validation

### Technologies

1. Golang
2. Aws lambda functions
4. AWS S3
5. Aws Dynamodb

### Next task

- Adding more validation in the build process like

*# Anything in before_script that returns a nonzero exit code will
*# flunk the build and immediately stop. It's sorta like having
*# set -e enabled in bash. 
before_script:
  - GO_FILES=$(find . -iname '*.go' -type f | grep -v /vendor/) # All the .go files, excluding vendor/
  - go get github.com/golang/lint/golint                        # Linter
  - go get honnef.co/go/tools/cmd/megacheck                     # Badass static analyzer/linter
  - go get github.com/fzipp/gocyclo

*# script always run to completion (set +e). All of these code checks are must haves
*# in a modern Go project.
script:
  - test -z $(gofmt -s -l $GO_FILES)         # Fail if a .go file hasn't been formatted with gofmt
  - go test -race -coverprofile=coverage.txt -covermode=atomic  # Run all the tests with the race detector enabled
  - go vet ./...                             # go vet is the official Go static analyzer
  - megacheck ./...                          # "go vet on steroids" + linter
  - gocyclo -over 19 $GO_FILES               # forbid code with huge functions
  - golint -set_exit_status $(go list ./...) # one last linter