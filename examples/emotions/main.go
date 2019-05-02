package main

import "log"

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}
