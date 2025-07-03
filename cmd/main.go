package main

import (
	"REST-API-pet-proj/internal"
	"fmt"
)

func main() {
	cfg := internal.InitConfigParser()
	fmt.Printf("%+v\n", cfg)
}
