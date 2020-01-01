Lambda example
============================

## build
```
$ go get -u github.com/aws/aws-lambda-go/lambda
$ go get -u github.com/aws/aws-sdk-go/aws
$ go get -u github.com/aws/aws-sdk-go/aws/credentials
$ go get -u github.com/aws/aws-sdk-go/aws/session
$ go get -u github.com/aws/aws-sdk-go/service/s3

GOOS=linux GOARCH=amd64 go build -o createBucket
```