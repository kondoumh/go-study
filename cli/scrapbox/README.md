Scrapbox module
==========================

run pages example

```
$ export GOBIN=`pwd`/_bin
$ go install github.com/kondoumh/go-study/cli/scrapbox/cmd/pages
$ _bin/pages
```

get page needs cookie

```
export COOKIE_NAME=cookie_name
export COOKIE_VALUE=fancy_cookie
go run cmd/page_detail/main.go
```