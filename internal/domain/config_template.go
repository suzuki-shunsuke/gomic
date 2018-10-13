package domain

// ConfigTpl is a configuration file template.
const ConfigTpl = `
---
items:
- src:
    package: github.com/suzuki-shunsuke/gomic/examples
    # file: examples/example.go
    interface: Hello
    name: HelloMock
  dest:
    package: example
    file: example/example_mock.go
`
