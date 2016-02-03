# go-base65536

Go library for encoding data into [base65536](https://github.com/ferno/base65536).

## Example

```go
package main

import (
	"fmt"
	"github.com/Nightbug/go-base65536"
)

func main() {
	fmt.Println(base65536.Marshal([]byte("hello world")))
}
```
