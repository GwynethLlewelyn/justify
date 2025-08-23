# justify
A Go package to justify text (left and right alignment), old-school typewriter style.

## Usage

```go
package main

import (
	"fmt"

	"github.com/GwynethLlewelyn/justify"
)

func main() {
	fmt.Println(justify.Justify(`
		
	`,
		80))	
}
```
