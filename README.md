# gomic

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/suzuki-shunsuke/gomic)
[![CircleCI](https://circleci.com/gh/suzuki-shunsuke/gomic.svg?style=svg)](https://circleci.com/gh/suzuki-shunsuke/gomic)
[![codecov](https://codecov.io/gh/suzuki-shunsuke/gomic/branch/master/graph/badge.svg)](https://codecov.io/gh/suzuki-shunsuke/gomic)
[![Go Report Card](https://goreportcard.com/badge/github.com/suzuki-shunsuke/gomic)](https://goreportcard.com/report/github.com/suzuki-shunsuke/gomic)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/gomic.svg)](https://github.com/suzuki-shunsuke/gomic)
[![GitHub tag](https://img.shields.io/github/tag/suzuki-shunsuke/gomic.svg)](https://github.com/suzuki-shunsuke/gomic/releases)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/gomic/master/LICENSE)

cli tool to generate golang's mock for test.

* [Features](#features)
* [Install](#install)
* [Getting Started](#getting-started)
* [Configuration](#configuration)
* [Note](#note)
* [Examples](https://github.com/suzuki-shunsuke/gomic/tree/master/examples)
* [Other Mocking Libraries](#other-mocking-libraries)
* [Change Log](https://github.com/suzuki-shunsuke/gomic/releases)

## Features

* Manage mocks with a configuration file, so it is easy to update mocks when interfaces are updated
* Generated mock is simple and generic. `EXPECT` and `RETURN` style isn't supported. You can set mock function freely.
* Provide simple fake implementation which only returns default values

## Install

gomic is written with Golang and binary is distributed at [release page](https://github.com/suzuki-shunsuke/gomic/releases), so installation is easy and no dependency is needed.

If you want to build yourself, run the following command.

```
$ go get -u github.com/suzuki-shunsuke/gomic/cmd/gomic
```

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

3. Edit the configuration file

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

```
$ gomic gen
```

**Note that `gen` command overwrites the existing file without confirmation.**

## Configuration

```yaml
---
# file path must be absolute or relative to the configuration file path.
# default is the default settings of each items. item's settings are preferred than default settings.
default:
#   vendor_dir: ""
#   interface_prefix: Mock
#   interface_suffix: Mock
items:
- src:
    # package or file or dir are required
    # package is a source package path.
    # file is a source file path.
    # dir is a source directory path.
    package: github.com/suzuki-shunsuke/gomic/examples
    # file: examples/example.go
    # source interface name. This is required.
    interface: Hello
    # generated mock name
    name: HelloMock
    # If name is not given, name is "{{interface_prefix}}{{interface}}{{interface_suffix}}".
    # If name is given, interface_prefix and interface_suffix are ignored.
    # interface_prefix: Mock
    # interface_suffix: Mock
    # vendor_dir is path of the parent directory of `vendor`.
    # vendor_dir should be absolute path or relative to configuration file's parent directory.
    # By default vendor_dir is configuration file's parent directory.
    #  vendor_dir: ""
  dest:
    # generated file's package name
    # If package is not set, gomic tries to get package name with file's parent directory path.
    package: example
    # output file path
    # file is required.
    # Currently it is not supported to output multiple mocks in the same file.
    # The parent directory must exist.
    file: example/example_mock.go
```

## Note

* `gen` command overwrites the existing file without confirmation
* Before run the `gen` command, packages which the interface depends on should be installed at GOPATH or vendor directory. gomic supports vendor directory.
* Currently it is not supported to output multiple mocks in the same file
* Generated code is not formatted. So we recommend to format them by `gofmt`.

## Examples

See [examples](https://github.com/suzuki-shunsuke/gomic/tree/master/examples) .

## Other Mocking Libraries

* https://github.com/avelino/awesome-go#testing
* https://github.com/golang/mock
* https://github.com/gojuno/minimock

## Change Log

Please see [Releases](https://github.com/suzuki-shunsuke/gomic/releases).

## Contributing

Please see [CONTRIBUTING.md](CONTRIBUTING.md) .

## License

[MIT](LICENSE)
