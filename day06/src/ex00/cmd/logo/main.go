package main

import (
	"log"

	"fortress-of-solitude/pkg/logo"
)

const name = "amazing_logo.png"

func main() {
	if err := logo.GenerateLogo(name); err != nil {
		log.Println(err)
	}
}
