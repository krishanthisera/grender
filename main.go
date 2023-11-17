package main

import (
	"flag"
	"fmt"

	"github.com/krishanthisera/grender/pilot"
)

func main() {
	fmt.Println("Grender is starting...")

	cfg := flag.String("config", "grender.yaml", "Path to grender config file")

	grenderConfig, err := pilot.GenerateConfig(*cfg)
	fmt.Print(grenderConfig)
	if err != nil {
		panic(err)
	}
	grenderConfig.Grender()
}
