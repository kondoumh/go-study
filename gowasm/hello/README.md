Hello wasm by golang
==================

[golang/go](https://github.com/golang/go/wiki/WebAssembly)

Using golang v1.12

## build
GOOS=js GOARCH=wasm go build -o main.wasm

## for browser
wasm_exec.js was copied from

https://github.com/golang/go/tree/dev.boringcrypto.go1.12/misc/wasm

```
$ go get -u github.com/shurcooL/goexec
$ goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'
```

## for Node.js

```
$ GOOS=js GOARCH=wasm go run -exec="$(go env GOROOT)/misc/wasm/go_js_wasm_exec" .
```