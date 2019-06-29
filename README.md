# gomic

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/suzuki-shunsuke/gomic)
[![Build Status](https://cloud.drone.io/api/badges/suzuki-shunsuke/gomic/status.svg)](https://cloud.drone.io/suzuki-shunsuke/gomic)
[![codecov](https://codecov.io/gh/suzuki-shunsuke/gomic/branch/master/graph/badge.svg)](https://codecov.io/gh/suzuki-shunsuke/gomic)
[![Go Report Card](https://goreportcard.com/badge/github.com/suzuki-shunsuke/gomic)](https://goreportcard.com/report/github.com/suzuki-shunsuke/gomic)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/gomic.svg)](https://github.com/suzuki-shunsuke/gomic)
[![GitHub tag](https://img.shields.io/github/tag/suzuki-shunsuke/gomic.svg)](https://github.com/suzuki-shunsuke/gomic/releases)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/gomic/master/LICENSE)

CLI tool to generate golang's mock for the test.

* [Features](#features)
* [Install](#install)
* [Getting Started](#getting-started)
* [Examples](https://github.com/suzuki-shunsuke/gomic/tree/master/examples)
* [Configuration](#configuration)
* [Note](#note)
* [Other Mocking Libraries](#other-mocking-libraries)
* [Change Log](https://github.com/suzuki-shunsuke/gomic/releases)

## Features

* Manage mocks with a configuration file, so it is easy to update mocks when interfaces are updated
* Simple and flexible API. Complicated API and DSL aren't provided. So the learning cost is low and the test code is easy to read

## Install

gomic is written with Golang and binary is distributed at the [release page](https://github.com/suzuki-shunsuke/gomic/releases), so installation is easy and no dependency is needed.

## Getting Started

1. Write an interface.

```go
type (
	OS interface {
		Getwd() (string, error)
		Mkdir(name string, perm os.FileMode) error
	}
)
```

2. Generate the configuration file.

```console
$ gomic init
```

3. Edit the configuration file.

```yaml
---
default:
  interface_suffix: Mock
items:
- src:
    file: os.go
    interface: OS
  dest:
    package: examples
    file: os_mock.go
- src:
    package: io
    interface: ReadCloser
  dest:
    package: examples
    file: readcloser_mock.go
```

4. Generate mocks.

```console
$ gomic gen
```

**Note that the `gen` command overwrites the existing file without confirmation.**

## Examples

See [examples](https://github.com/suzuki-shunsuke/gomic/tree/master/examples).

## Configuration

```yaml
---
# the file path must be absolute or relative to the configuration file path.
# `default` is the default settings of each item. item's settings are preferred than default settings.
default:
#   vendor_dir: ""
#   interface_prefix: Mock
#   interface_suffix: Mock
items:
- src:
    # `package` or `file` or `dir` are required
    # `package` is the source package path.
    # `file` is the source file path.
    # `dir` is the source directory path.
    package: github.com/suzuki-shunsuke/gomic/examples
    # file: examples/example.go
    # source interface name. This is required.
    interface: Hello
    # generated mock name
    name: HelloMock
    # If `name` is not given, name is "{{interface_prefix}}{{interface}}{{interface_suffix}}".
    # If `name` is given, `interface_prefix` and `interface_suffix` are ignored.
    # interface_prefix: Mock
    # interface_suffix: Mock
    # `vendor_dir` is the path of the parent directory of `vendor`.
    # `vendor_dir` should be absolute path or relative to configuration file's parent directory.
    # By default `vendor_dir` is the configuration file's parent directory.
    #  vendor_dir: ""
  dest:
    # generated file's package name
    # If `package` is not set, gomic tries to get the package name with the file's parent directory path.
    package: example
    # output file path
    # `file` is required.
    # Currently it is not supported to output multiple mocks in the same file.
    # The parent directory must exist.
    file: example/example_mock.go
```

## Note

* `gen` command overwrites the existing file without confirmation
* Before running the `gen` command, packages which the interface depends on should be installed at GOPATH or vendor directory. gomic supports vendor directory.
* Currently, it is not supported to output multiple mocks in the same file
* The generated code is not formatted. So we recommend to format them by `gofmt`.

## Other Mocking Libraries

* https://github.com/avelino/awesome-go#testing
* https://github.com/golang/mock
* https://github.com/gojuno/minimock

## Troble shooting

### go: cannot find GOROOT directory

If you encounter the following error when you run `gomic gen`,
try to set the environment variable `GOROOT`.

```
go/build: importGo %s: exit status 2
go: cannot find GOROOT directory: %s
```

```console
$ go version
go version go1.12.6 darwin/amd64
$ gomic -v
gomic version 0.5.6
$ echo $GOROOT  # GOROOT isn't defined

$ cat .gomic.yml
---
items:
- src:
    package: io
    interface: Writer
  dest:
    package: mock
    file: mock/io_writer.go
$ gomic gen
go/build: importGo io: exit status 2
go: cannot find GOROOT directory: /usr/local/go
```

We don't know the root cause, but we can resolve this issue in this way.

```console
$ go env GOROOT
/usr/local/Cellar/go/1.12.6/libexec
GOROOT=`go env GOROOT` gomic gen
```

## Change Log

Please see the [Releases](https://github.com/suzuki-shunsuke/gomic/releases).

## Contributing

Please see the [CONTRIBUTING.md](CONTRIBUTING.md).

## License

[MIT](LICENSE)
