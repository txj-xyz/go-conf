package main

import (
	"fmt"

	"github.com/txj-xyz/go-conf/config"
)

func main() {
	fmt.Println("Hello from main")
	cfg, err := config.Load()
	if err != nil {
		fmt.Errorf("error loading config: %v", err)
	}

	// fmt.Printf("config loaded from main: %v", cfg)
	cfg.PrettyPrint()

}
