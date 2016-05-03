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

## License

The [BSD 3-Clause license][bsd], the same as the [Go language][golic]. Cascadia's license is [here][caslic].

[jquery]: http://jquery.com/
[go]: http://golang.org/
[cascadia]: https://github.com/andybalholm/cascadia
[bsd]: http://opensource.org/licenses/BSD-3-Clause
[golic]: http://golang.org/LICENSE
[caslic]: https://github.com/andybalholm/cascadia/blob/master/LICENSE
[doc]: http://godoc.org/github.com/PuerkitoBio/goquery
[index]: http://api.jquery.com/index/
[gonet]: https://github.com/golang/net/
[html]: http://godoc.org/golang.org/x/net/html
[wiki]: https://github.com/PuerkitoBio/goquery/wiki/Tips-and-tricks
[thatguystone]: https://github.com/thatguystone
[piotr]: https://github.com/piotrkowalczuk
