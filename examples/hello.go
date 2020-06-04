package main

import (
	"fmt"
	client "interview-accountapi/internal/http"
)

func main() {
	err := client.Health()
	fmt.Printf("hello world %v\n", err)
}
