[![Build Status](https://travis-ci.org/alexrv11/lambda-api-taxi-friend.svg?branch=master)](https://travis-ci.org/alexrv11/lambda-api-taxi-friend)
[![codecov](https://codecov.io/gh/alexrv11/lambda-api-taxi-friend/branch/master/graph/badge.svg)](https://codecov.io/gh/alexrv11/lambda-api-taxi-friend)

# Taxifriend server

The taxifriend server is an application to provides taxi services, monitoring taxi location, provide information in a near arear, register new drivers

### Technologies

1. Golang
2. Aws lambda functions
4. AWS S3
5. Aws Dynamodb
6. AWS API Gateway
7. AWS Serverless Application Model
8. AWS Cognito

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
  
  
## lambda-api-taxi-friend

This is a sample template for lambda-api-taxi-friend - Below is a brief explanation of what we have generated for you:

```bash
.
├── Makefile                    <-- Make to automate build
├── README.md                   <-- This instructions file
├── hello-world                 <-- Source code for a lambda function
│   ├── main.go                 <-- Lambda function code
│   └── main_test.go            <-- Unit tests
└── template.yaml
```

## Requirements

* AWS CLI already configured with Administrator permission
* [Docker installed](https://www.docker.com/community-edition)
* [Golang](https://golang.org)

## Setup process

### Installing dependencies

In this example we use the built-in `go get` and the only dependency we need is AWS Lambda Go SDK:

```shell
go get -u github.com/aws/aws-lambda-go/...
```

**NOTE:** As you change your application code as well as dependencies during development, you might want to research how to handle dependencies in Golang at scale.

### Building

Golang is a statically compiled language, meaning that in order to run it you have to build the executable target.

You can issue the following command in a shell to build it:

```shell
GOOS=linux GOARCH=amd64 go build -o hello-world/hello-world ./hello-world
```

**NOTE**: If you're not building the function on a Linux machine, you will need to specify the `GOOS` and `GOARCH` environment variables, this allows Golang to build your function for another system architecture and ensure compatibility.

### Local development

**Invoking function locally through local API Gateway**

```bash
sam local start-api
```

If the previous command ran successfully you should now be able to hit the following local endpoint to invoke your function `http://localhost:3000/hello`

**SAM CLI** is used to emulate both Lambda and API Gateway locally and uses our `template.yaml` to understand how to bootstrap this environment (runtime, where the source code is, etc.) - The following excerpt is what the CLI will read in order to initialize an API and its routes:

```yaml
...
Events:
    HelloWorld:
        Type: Api # More info about API Event Source: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#api
        Properties:
            Path: /hello
            Method: get
```

## Packaging and deployment

AWS Lambda Golang runtime requires a flat folder with the executable generated on build step. SAM will use `CodeUri` property to know where to look up for the application:

```yaml
...
    FirstFunction:
        Type: AWS::Serverless::Function
        Properties:
            CodeUri: hello_world/
            ...
```

To deploy your application for the first time, run the following in your shell:

```bash
sam deploy --guided
```

The command will package and deploy your application to AWS, with a series of prompts:

* **Stack Name**: The name of the stack to deploy to CloudFormation. This should be unique to your account and region, and a good starting point would be something matching your project name.
* **AWS Region**: The AWS region you want to deploy your app to.
* **Confirm changes before deploy**: If set to yes, any change sets will be shown to you before execution for manual review. If set to no, the AWS SAM CLI will automatically deploy application changes.
* **Allow SAM CLI IAM role creation**: Many AWS SAM templates, including this example, create AWS IAM roles required for the AWS Lambda function(s) included to access AWS services. By default, these are scoped down to minimum required permissions. To deploy an AWS CloudFormation stack which creates or modified IAM roles, the `CAPABILITY_IAM` value for `capabilities` must be provided. If permission isn't provided through this prompt, to deploy this example you must explicitly pass `--capabilities CAPABILITY_IAM` to the `sam deploy` command.
* **Save arguments to samconfig.toml**: If set to yes, your choices will be saved to a configuration file inside the project, so that in the future you can just re-run `sam deploy` without parameters to deploy changes to your application.

You can find your API Gateway Endpoint URL in the output values displayed after deployment.

### Testing

We use `testing` package that is built-in in Golang and you can simply run the following command to run our tests:

```shell
go test -v ./...
```
# Appendix

### Golang installation

Please ensure Go 1.x (where 'x' is the latest version) is installed as per the instructions on the official golang website: https://golang.org/doc/install

