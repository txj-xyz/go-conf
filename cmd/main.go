package main

import (
	"fmt"

	"github.com/txj-xyz/go-conf/config"
)

func main() {
	fmt.Println("Hello from main")
	config.Load()

}
