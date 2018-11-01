package domain

// ConfigTpl is a configuration file template.
const ConfigTpl = `
---
default:
  # interface_prefix: Mock
  interface_suffix: Mock
items:
- src:
    package: os
    interface: FileInfo
    # name: FileInfoMock
  dest:
    package: test
    file: test/fileinfo_mock.go
`
