package main

import (
	"fmt"

	middleware "github.com/damilarelana/goSystemDesign/middleware"
)

var appName = "accountService"

func main() {
	fmt.Printf("Starting %v\n", appName)
	middleware.InitializeWebServer("8080")
}
