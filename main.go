package main

import (
	"fmt"

	"github.com/krishanthisera/grender/pilot"
)

func main() {
	fmt.Println("Grender is starting...")
	grenderConfig, err := pilot.GenerateConfig("grender.yaml")
	fmt.Print(grenderConfig)
	if err != nil {
		panic(err)
	}
	grenderConfig.Grender()
}
