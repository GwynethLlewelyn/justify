# justify
![Justify Gopher Logo](assets/justify-gopher-logo-128x128.png)  
A Go package to justify text (left and right alignment), old-school typewriter style.

## Usage

```go
package main

import (
	"fmt"

	"github.com/GwynethLlewelyn/justify"
)

func main() {
	fmt.Println(justify.Justify(`Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum`, 40))	
}
```

produces the following output:

```
Lorem  ipsum dolor sit amet, consectetur
adipiscing  elit,  sed do eiusmod tempor
incididunt  ut  labore  et  dolore magna
aliqua.  Ut  enim  ad minim veniam, quis
nostrud   exercitation  ullamco  laboris
nisi ut aliquip ex ea commodo consequat.
Duis  aute  irure dolor in reprehenderit
in voluptate velit esse cillum dolore eu
fugiat  nulla  pariatur.  Excepteur sint
occaecat cupidatat non proident, sunt in
culpa  qui  officia deserunt mollit anim
id est laborum.
```

## Variables

To avoid outputting the last newline (whenever calling this function from somewhere that already properly inserts a newline at the end), just set `justify.EOL = false`. Not goroutine-safe. 


[![Go](https://github.com/GwynethLlewelyn/justify/actions/workflows/go.yml/badge.svg)](https://github.com/GwynethLlewelyn/justify/actions/workflows/go.yml)