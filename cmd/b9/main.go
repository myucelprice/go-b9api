package main

import (
	"encoding/json"
	"fmt"

	"github.com/myucelprice/go-b9api"
)

func main() {
	datas, _ := b9.GetHD(2)
	for _, data := range datas {
		js, _ := json.MarshalIndent(data, "", "\t")
		fmt.Println(string(js))
	}
}
