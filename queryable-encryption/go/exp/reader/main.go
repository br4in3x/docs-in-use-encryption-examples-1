package main

import (
	"goCsfle/local"
	"log"
)

func main() {
	err := local.MakeKey()
	if err != nil {
		log.Fatal(err)
	}
	err = local.Insert()
	if err != nil {
		log.Fatal(err)
	}

}
