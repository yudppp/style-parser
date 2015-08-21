[![License](http://img.shields.io/:license-mit-brightgreen.svg?style=flat-square)](http://yudppp.mit-license.org)

# style-parser

parse style


## Quick Start

```go
package main

import (
  "fmt"
  "github.com/yudppp/style-parser"
)

func main() {
  parser := Parser{}
  styleStr := `color:blue; line-height:1.5;`
  set := parser.ParseFromString(styleStr)
  fmt.Println(set)
  // output: map[color:blue line-height:1.5]
}
```
