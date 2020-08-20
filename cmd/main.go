package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/seymores/carouscrapy"
)

func main() {
	check := len(os.Args)
	if check < 2 {
		fmt.Println("Usage: carouscapy <url>")
		os.Exit(3)
	}

	url := flag.String("url", "", "Item URL")

	flag.Parse()

	// url := "https://sg.carousell.com/p/ricoh-gr-iii-1028411869"

	carouscrapy.LoadByURL(*url)
}
