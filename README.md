# nestedmap [![Go Reference](https://pkg.go.dev/badge/github.com/ari1021/nestedmap.svg)](https://pkg.go.dev/github.com/ari1021/nestedmap)

`nestedmap` is a nested map for Go.

`nestedmap` prevents `panic: assignment to entry in nil map` and we don't have to worry about map initialization.

## requirement

`nestedmap` requires Go 1.18+.

## how to use

### install

```
$ go get github.com/ari1021/nestedmap
```

### usage

```go
package main

import (
	"fmt"

	"github.com/ari1021/nestedmap"
)

func main() {
	// Initialize nestedmap with type argument
	// map[string]map[int]bool
	nm := nestedmap.NewNestedMap[string, int, bool]()

	// Set value
	nm.Set("gopher", 117, false)
	nm.Set("gopher", 118, true)
	// map[gopher:map[117:false 118:true]]]

	// Get outer value. The type is map[int]bool
	if m, ok := nm.GetOuter("gopher"); ok {
		fmt.Println(m) // output: map[117:true 118:true]
	}

	// Get inner value. The type is bool
	if v, ok := nm.GetInner("gopher", 118); ok {
		fmt.Println(v) // output: true
	}

	// When we try to get non-existing data, ok is false
	if v, ok := nm.GetInner("gooopher", 118); ok {
		fmt.Println(v) // not called
	}
}
```