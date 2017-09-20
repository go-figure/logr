# logr
[![Build Status](http://img.shields.io/travis/go-figure/logr/master.svg?style=flat-square)](https://travis-ci.org/go-figure/logr)
[![Coverage Status](http://img.shields.io/coveralls/go-figure/logr/master.svg?style=flat-square)](https://coveralls.io/r/go-figure/logr)
[![GoDoc](http://img.shields.io/badge/api-Godoc-blue.svg?style=flat-square)](https://godoc.org/github.com/go-figure/logr)
[![License: MIT](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](http://opensource.org/licenses/MIT)

`logr` aims to provide powerful, composable and extendable logging for Go
applications and libraries.


## Goals
  - Provide a way of logging as much data as possible from the application and
    its dependencies
  - Add structure to the logging facilities needed throughout the application
  - Simplify inserting context data
  - Easy extensibility trough the use of few simple interfaces
  - Provide easy access to logging trough the `context` package
  - Ability to sink different kind of logs to different sinks


## Status
Tests and mocks are provided for most of the code base, and `logr` is actively
being used in production in serveral projects.

[Documentation](https://godoc.org/github.com/go-figure/logr) is currently lacking.


## Contributing
Contributing is always welcome and appreciated! Pull requests and issues will
be processed as fast as possible!


## License
Copyright (c) 2017 George-Cristian Jiglau

This project is licensed under the [MIT license](http://opensource.org/licenses/MIT).
See the LICENSE file for more details.
