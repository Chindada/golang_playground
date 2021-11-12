package main

import (
	"fmt"
	"os"
)

func main() {
	appdata := os.Getenv("APPDATA")
	fmt.Println(appdata)
}
