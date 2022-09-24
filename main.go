package main

import (
	"flag"
	"log"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		log.Fatal("コマンドライン引数に line か scrape を指定してください")
	}

	switch args[0] {
	case "line":
		lineNotify()
	case "scrape":
		scrape()
	default:
		log.Fatal("コマンドライン引数に line か　scrape を指定してください")
	}
}
