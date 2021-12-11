[![Sourcegraph](https://sourcegraph.com/github.com/tint/env/-/badge.svg?style=flat-square)](https://sourcegraph.com/github.com/tint/env?badge)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/tint/env)
[![Go Report Card](https://goreportcard.com/badge/github.com/tint/env?style=flat-square)](https://goreportcard.com/report/github.com/tint/env)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/tint/env/master/LICENSE)

## Installation

```sh
go get github.com/tint/env
```

### 初始化

自动检测应用目录

```go
env.Setup()
```

自定义启动目录

```go
env.Setup("path/to/directory")
```

### 获取值

```go
  env.Lookup(name string) (val string, exists bool)
  env.Exists(name string) bool
  env.Unknown(name string, value ...interface{}) interface{}
  env.String(name string, value ...string) string
  env.Bytes(name string, value ...[]byte) []byte
  env.Int(name string, value ...int) int
  env.Duration(name string, value ...time.Duration) time.Duration
  env.Bool(name string, value ...bool) bool
  env.Map(prefix string) map[string]string
  env.Where(filter func(name string, value string) bool) map[string]string
  env.Environ() []string
```
