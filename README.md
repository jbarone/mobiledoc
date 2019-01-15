## Mobiledoc Renderer

[![MIT License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](/LICENSE)

This is a library for rendering the [Mobiledoc
format](https://github.com/bustlelabs/mobiledoc-kit/blob/master/MOBILEDOC.md)
used by [Mobiledoc-Kit](https://github.com/bustlelabs/mobiledoc-kit).

Currently this library only supports rendering to Markdown.

## Motivation

This project was created to help with the conversion of Ghost Blog exports to
Hugo Blog format, using the tool
[ghostToHugo](http://github.com/jbarone/ghostToHugo). Starting in version 2.0.0
Ghost switched from using Markdown to using Mobiledoc as the underlying format
of the post content. This made the conversion tool requirements more complicated
since understanding Mobiledoc was now necessary.

## Build status

[![GitHub release](https://img.shields.io/github/release/jbarone/mobiledoc.svg)](https://github.com/jbarone/mobiledoc/releases/latest)
[![Build Status](https://travis-ci.org/jbarone/mobiledoc.svg?branch=master)](https://travis-ci.org/jbarone/mobiledoc)
[![Go Report Card](https://goreportcard.com/badge/github.com/jbarone/mobiledoc)](https://goreportcard.com/report/github.com/jbarone/mobiledoc)

## Code Example

Show what the library does as concisely as possible, developers should be able
to figure out **how** your project solves their problem by looking at the code
example. Make sure the API you are showing off is obvious, and that your code is
short and concise.

## Installation

Mobiledoc is a Go package library and is intended to be used by other tools. To
make use of this library, you just need to use the go tools to fetch the library
and then import into your code.

```
$ go get -u github.com/jbarone/mobiledoc
```

```go
package main

import "github.com/jbarone/mobiledoc"
```

## API Reference

[GoDocs](https://godoc.org/github.com/jbarone/mobiledoc)

## Tests

Tests are written using the standard go testing package.

```
$ go test
```

## Credits

Much credit is given to
[mobiledoc-markdown-renderer](https://github.com/yuloh/mobiledoc-markdown-renderer),
from which much inspiration for this code was taken.

## License

Released under the [MIT](./LICENSE) license.
