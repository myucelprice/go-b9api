# go-b9api

## exsample

```
package main

import (
	"encoding/json"
	"fmt"

	"github.com/ieee0824/b9"
)

func main() {
	datas, _ := b9.GetHD(2)
	for _, data := range datas {
		js, _ := json.MarshalIndent(data, "", "\t")
		fmt.Println(string(js))
	}
}
```
