## go-compressed

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/ferhatelmas/go-compressed)
[![Build Status](https://travis-ci.org/ferhatelmas/go-compressed.png?branch=master)](https://travis-ci.org/ferhatelmas/go-compressed)

> Check if a filepath is a compressed file.

### Install

```
go get github.com/ferhatelmas/go-compressed
```

### Usage

```go
import "github.com/ferhatelmas/go-compressed"

compressed.Is("src/unicorn.rar")
//=> true

compressed.Is("src/unicorn.txt")
//=> false
```

### License

MIT © [ferhat elmas](http://ferhatelmas.com)
