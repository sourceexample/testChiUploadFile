package main

import (
	"chiFile/modHttp"
	"fmt"
)

func main() {
	err := modHttp.Chi_Initialize()
	if err != nil {
		fmt.Println(err)
	}
}
